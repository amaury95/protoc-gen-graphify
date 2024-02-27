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
	g.P(`
		// parseFloat ...
		func parseFloat(s string) float64 {
			var res, decFactor float64
			var neg, decSeen bool
			for _, char := range s {
				switch char {
				case '-':
					if res != 0 || neg {
						return 0
					}
					neg = true
				case '.':
					if decSeen {
						return 0
					}
					decSeen = true
					decFactor = 0.1
				default:
					if char < '0' || char > '9' {
						return 0
					}
					digit := float64(char - '0')
					if decSeen {
						res = res + digit*decFactor
						decFactor *= 0.1
					} else {
						res = res*10 + digit
					}
				}
			}
			if neg {
				res = -res
			}
			return res
		}
	`)
	g.P()
	g.P(`
		// makeSlice ... 
		func makeSlice(ptr interface{}, size int) {
			ptrVal := reflect.ValueOf(ptr)
			if ptrVal.Kind() == reflect.Ptr && ptrVal.Elem().Kind() == reflect.Slice {
				_n := reflect.MakeSlice(ptrVal.Elem().Type(), size, size)
				ptrVal.Elem().Set(_n)
			}
		}
	`)
	g.P()
	g.P(`
		// makeMap ... 
		func makeMap(ptr interface{}) {
			mapVal := reflect.ValueOf(ptr)
			if mapVal.Kind() == reflect.Ptr && mapVal.Elem().Kind() == reflect.Map {
				newMap := reflect.MakeMap(mapVal.Elem().Type())
				mapVal.Elem().Set(newMap)
			}
		}
	`)
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
			if field.Desc.IsSynthetic() {
				continue
			}
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
		g.P("if val , ok := _opt[\"", oneofField.GoName, "\"].(map[string]interface{}); ok {")
		g.P("field := new(", oneofField.Message.GoIdent.GoName, ")")
		g.P("field.LoadMap(val)")
		g.P(recipient, assign, "&", oneofField.GoIdent, "{", oneofField.GoName, ":field}")
		g.P("}")
	}
	g.P("}")
}

func fetchField(g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	fetchField_Exportable(false, g, field, recipient, assign, identifier...)
}

func fetchAndExportField(g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	fetchField_Exportable(true, g, field, recipient, assign, identifier...)
}

func fetchField_Exportable(export bool, g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	switch {
	case field.Desc.IsList():
		g.P(join("if list , ok := ", identifier, ".([]interface{}); ok {")...)
		g.P("makeSlice(&", recipient, ", len(list))")
		g.P("for index, item := range list {")
		fetchField(g, ignoreType(field), recipient+"[index]", "=", "item")
		g.P("}")
		g.P("}")
		return
	case field.Desc.IsMap():
		keyField := field.Message.Fields[0]
		valField := field.Message.Fields[1]
		g.P(join("if values, ok := ", identifier, ".(map[string]interface{}); ok {")...)
		g.P("makeMap(&", recipient, ")")
		g.P("for key, value := range values {")
		if keyField.Desc.Kind() == protoreflect.StringKind {
			fetchField(g, valField, recipient+"[key]", "=", "value")
		} else {
			g.P("var tmp interface{} = parseFloat(key)")
			fetchAndExportField(g, keyField, "parsedKey", "=", "tmp")
			fetchField(g, valField, recipient+"[parsedKey]", "=", "value")
		}
		g.P("}")
		g.P("}")
		return
	}

	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		assignMessage(export, g, field, recipient, assign, identifier...)
	case protoreflect.BoolKind:
		assignField(export, g, field, "bool", recipient, assign, identifier...)
	case protoreflect.EnumKind:
		assignEnum(export, g, field, recipient, assign, identifier...)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		parseField(export, g, field, "float64", "int32", recipient, assign, identifier...)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		parseField(export, g, field, "float64", "uint32", recipient, assign, identifier...)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		parseField(export, g, field, "float64", "int64", recipient, assign, identifier...)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		parseField(export, g, field, "float64", "uint64", recipient, assign, identifier...)
	case protoreflect.FloatKind:
		parseField(export, g, field, "float64", "float32", recipient, assign, identifier...)
	case protoreflect.DoubleKind:
		assignField(export, g, field, "float64", recipient, assign, identifier...)
	case protoreflect.StringKind:
		assignField(export, g, field, "string", recipient, assign, identifier...)
	case protoreflect.BytesKind:
		parseField(export, g, field, "string", "[]byte", recipient, assign, identifier...)
	}
}

func assignMessage(export bool, g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	g.P(join("if val , ok := ", identifier, ".(map[string]interface{}); ok {")...)
	g.P("field := new(", field.Message.GoIdent.GoName, ")")
	g.P("field.LoadMap(val)")
	g.P(recipient, assign, "field")
	g.P("}")
}

func assignEnum(export bool, g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	parseField(true, g, field, "float64", "int32", "tmp", " = ", identifier...)
	g.P(recipient, assign, g.QualifiedGoIdent(field.Enum.GoIdent), "(tmp)")
}

func assignField(export bool, g *protogen.GeneratedFile, field *protogen.Field, typeTo string, recipient, assign string, identifier ...interface{}) {
	if export {
		g.P("var ", recipient, " ", typeTo)
	}
	if field.Desc.HasPresence() {
		g.P(join("if val, ok := ", identifier, ".(", typeTo, "); ok {")...)
		g.P(recipient, assign, "&val")
		g.P("}")
	} else {
		g.P(join("if val, ok := ", identifier, ".(", typeTo, "); ok {")...)
		g.P(recipient, assign, "val")
		g.P("}")
	}
}

func parseField(export bool, g *protogen.GeneratedFile, field *protogen.Field, typeFrom, typeTo string, recipient, assign string, identifier ...interface{}) {
	if export {
		g.P("var ", recipient, " ", typeTo)
	}
	if field.Desc.HasPresence() {
		g.P(join("if val, ok := ", identifier, ".(", typeFrom, "); ok {")...)
		g.P("tmp := ", typeTo, "(val)")
		g.P(recipient, assign, "&tmp")
		g.P("}")
	} else {
		g.P(join("if val, ok := ", identifier, ".(", typeFrom, "); ok {")...)
		g.P(recipient, assign, typeTo, "(val)")
		g.P("}")
	}
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
