// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go binary is a protoc plugin to generate Go code for
// both proto2 and proto3 versions of the protocol buffer language.
//
// For more information about the usage of this plugin, see:
// https://protobuf.dev/reference/go/go-generated.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/amaury95/protoc-gen-go-tag/internal/version"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const genGoDocURL = "https://protobuf.dev/reference/go/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\n" +
				"See " + grpcDocURL + " for more information")
		}
		for _, f := range gen.Files {
			if f.Generate {
				messages := walkMessages(f.Messages)
				updateMessageNames(messages...)
				g := gengo.GenerateFile(gen, f)
				exposeMapBuilders(g, f, messages...)
				setHelpers(g)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}

func walkMessages(messages []*protogen.Message) []*protogen.Message {
	var result []*protogen.Message
	for _, message := range messages {
		result = append(result, message)
		result = append(result, walkMessages(message.Messages)...)
	}
	return result
}

func setHelpers(g *protogen.GeneratedFile) {
	g.P()
	g.P("func parseFloat(s string) float64 { var _r, _df float64 ;var _n, _ds bool;for _, char := range s { switch char { case '-': if _r != 0 || _n { return 0};_n = true; case '.': if _ds { return 0 } ;_ds = true ;	_df = 0.1 ;default: if char < '0' || char > '9' { return 0 } ; digit := float64(char - '0') ; if _ds { _r = _r + digit*_df ; _df *= 0.1 } else { _r = _r*10 + digit } } } ;if _n { _r = -_r }; return _r }")
	g.P()
	g.P("func makeSlice(_ptr interface{}, _s int) { _val := reflect.ValueOf(_ptr);if _val.Kind() == reflect.Ptr && _val.Elem().Kind() == reflect.Slice { _n := reflect.MakeSlice(_val.Elem().Type(), _s, _s); _val.Elem().Set(_n) } }")
	g.P()
}

func updateMessageNames(messages ...*protogen.Message) {
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.Desc.HasJSONName() {
				field.Desc = overrideJsonName{name: field.Desc.JSONName(), FieldDescriptor: field.Desc}
			}
		}
	}
}

func exposeMapBuilders(g *protogen.GeneratedFile, f *protogen.File, messages ...*protogen.Message) {
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}
		g.P()
		g.P("// LoadMap ...")
		g.P("func (o *", message.GoIdent, ") LoadMap(values map[string]interface{}) {")
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}
			fetchField(g, field, "o."+field.GoName, " = ", "values[\"", field.Desc.Name(), "\"]")
		}
		for _, field := range message.Oneofs {
			fetchOneof(g, field, "o."+field.GoName, " = ", "values[\"", field.GoName, "\"]")
		}
		g.P("}")
	}
}

func fetchOneof(g *protogen.GeneratedFile, field *protogen.Oneof, recipient, assign string, identifier ...interface{}) {
	g.P(join("if _opt, ok := ", identifier, ".(map[string]interface{}); ok {")...)
	for _, oneofField := range field.Fields {
		if oneofField.Message == nil {
			continue
		}
		g.P("if _val , ok := _opt[\"", oneofField.GoName, "\"].(map[string]interface{}); ok {")
		g.P("field := new(", oneofField.Message.GoIdent.GoName, ")")
		g.P("field.LoadMap(_val)")
		g.P(recipient, assign, "&", oneofField.GoIdent, "{", oneofField.GoName, ":field}")
		g.P("}")
	}
	g.P("}")
}

func fetchField(g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	switch {
	case field.Desc.IsList():
		g.P(join("if _list , ok := ", identifier, ".([]interface{}); ok {")...)
		g.P("makeSlice(&", recipient, ", len(_list))")
		g.P("for _i, _item := range _list {")
		fetchField(g, ignoreType(field), recipient+"[_i]", "=", "_item")
		g.P("}")
		g.P("}")
		return
	case field.Desc.IsMap():
		// keyType, _, _ := fieldGoType(g, f, field.Message.Fields[0])
		// valType, _, _ := fieldGoType(g, f, field.Message.Fields[1])
		// g.P(join("if _ , ok := ", identifier, ".(map[string]interface{}); ok {")...)

		// g.P("}")
		return
	}

	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P(join("if _val , ok := ", identifier, ".(map[string]interface{}); ok {")...)
		g.P("field := new(", field.Message.GoIdent.GoName, ")")
		g.P("field.LoadMap(_val)")
		g.P(recipient, assign, "field")
		g.P("}")
	case protoreflect.BoolKind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(bool); ok {")...)
			g.P(recipient, assign, "&_val")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(bool); ok {")...)
			g.P(recipient, assign, "_val")
			g.P("}")
		}
	case protoreflect.EnumKind:
		// goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(join("if ok {")...)
			g.P("_d := int32(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "int32(_val)")
			g.P("}")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P("_d := uint32(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "uint32(_val)")
			g.P("}")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P("_d := int64(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "int64(_val)")
			g.P("}")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P("_d := uint64(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "uint64(_val)")
			g.P("}")
		}
	case protoreflect.FloatKind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P("_d := float32(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "float32(_val)")
			g.P("}")
		}
	case protoreflect.DoubleKind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "&_val")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(float64); ok {")...)
			g.P(recipient, assign, "_val")
			g.P("}")
		}
	case protoreflect.StringKind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(string); ok {")...)
			g.P(recipient, assign, "&_val")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(string); ok {")...)
			g.P(recipient, assign, "_val")
			g.P("}")
		}
	case protoreflect.BytesKind:
		if field.Desc.HasPresence() {
			g.P(join("if _val, ok := ", identifier, ".(string); ok {")...)
			g.P("_d := []byte(_val)")
			g.P(recipient, assign, "&_d")
			g.P("}")
		} else {
			g.P(join("if _val, ok := ", identifier, ".(string); ok {")...)
			g.P(recipient, assign, "[]byte(_val)")
			g.P("}")
		}
	}
}

// fieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGoType(g *protogen.GeneratedFile, f *protogen.File, field *protogen.Field) (goType string, parseType *string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", nil, false
	}

	float64T := "float64"

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
		parseType = &float64T
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
		parseType = &float64T
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
		parseType = &float64T
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
		parseType = &float64T
	case protoreflect.FloatKind:
		goType = "float32"
		parseType = &float64T
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
		// case protoreflect.MessageKind, protoreflect.GroupKind:
		// 	goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		// 	pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, nil, false
	case field.Desc.IsMap():
		keyType, _, _ := fieldGoType(g, f, field.Message.Fields[0])
		valType, _, _ := fieldGoType(g, f, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), nil, false
	}
	return goType, parseType, pointer
}

func join(parts ...interface{}) (result []interface{}) {
	for _, part := range parts {
		partValue := reflect.ValueOf(part)

		if partValue.Kind() == reflect.Slice {
			// If part is a slice, append its elements to the result
			for i := 0; i < partValue.Len(); i++ {
				result = append(result, partValue.Index(i).Interface())
			}
		} else {
			// If part is not a slice, append it directly to the result
			result = append(result, part)
		}
	}
	return result
}

func ignoreType(f *protogen.Field) *protogen.Field {
	f.Desc = &ignoreDesc{FieldDescriptor: f.Desc}
	return f
}

type ignoreDesc struct {
	protoreflect.FieldDescriptor
}

func (*ignoreDesc) IsList() bool {
	return false
}
