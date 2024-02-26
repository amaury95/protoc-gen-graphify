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
				"See " + grpcDocURL + " for more information.")
		}
		for _, f := range gen.Files {
			if f.Generate {
				for _, message := range f.Messages {
					for _, field := range message.Fields {
						if field.Desc.HasJSONName() {
							field.Desc = overrideJsonName{name: field.Desc.JSONName(), FieldDescriptor: field.Desc}
						}
					}
				}

				g := gengo.GenerateFile(gen, f)
				// exposeOneofWrappers(g, walkMessages(f.Messages)...)
				exposeMapBuilders(g, f, walkMessages(f.Messages)...)
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

func exposeMapBuilders(g *protogen.GeneratedFile, f *protogen.File, messages ...*protogen.Message) {
	for _, message := range messages {
		g.P()
		g.P("// LoadMap ...")
		g.P("func (e *", message.GoIdent, ") LoadMap(m map[string]interface{}) {")
		for _, field := range message.Fields {
			g.P("// ", field.Desc.Name())
			if field.Oneof != nil {
				continue
			}

			switch field.Desc.Kind() {
			case protoreflect.MessageKind:
				g.P("if _val , ok := m[\"", field.Desc.Name(), "\"].(map[string]interface{}); ok {")
				g.P("field := new(", field.Message.GoIdent.GoName, ")")
				g.P("field.LoadMap(_val)")
				g.P("e.", field.GoName, " = field")
				g.P("}")
			case protoreflect.BoolKind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(bool); ok {")
					g.P("e." + field.GoName + " = &_val")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(bool); ok {")
					g.P("e." + field.GoName + " = _val")
					g.P("}")
				}
			case protoreflect.EnumKind:
				// goType = g.QualifiedGoIdent(field.Enum.GoIdent)
			case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &(int32(_val))")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = int32(_val)")
					g.P("}")
				}
			case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &(uint32(_val))")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = uint32(_val)")
					g.P("}")
				}
			case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &(int64(_val))")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = int64(_val)")
					g.P("}")
				}
			case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &(uint64(_val))")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = uint64(_val)")
					g.P("}")
				}
			case protoreflect.FloatKind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &(float32(_val))")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = float32(_val)")
					g.P("}")
				}
			case protoreflect.DoubleKind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = &_val")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(float64); ok {")
					g.P("e." + field.GoName + " = _val")
					g.P("}")
				}
			case protoreflect.StringKind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(string); ok {")
					g.P("e." + field.GoName + " = &_val")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].(string); ok {")
					g.P("e." + field.GoName + " = _val")
					g.P("}")
				}
			case protoreflect.BytesKind:
				if field.Desc.HasPresence() {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].([]byte); ok {")
					g.P("e." + field.GoName + " = &_val")
					g.P("}")
				} else {
					g.P("if _val, ok := m[\"", field.Desc.Name(), "\"].([]byte); ok {")
					g.P("e." + field.GoName + " = _val")
					g.P("}")
				}
			}
		}

		for _, oneof := range message.Oneofs {
			if oneof.Desc.IsSynthetic() {
				continue
			}
			g.P("if _opt, ok := m[\"" + oneof.GoName + "\"].(map[string]interface{}); ok {")
			for _, oneofField := range oneof.Fields {
				if oneofField.Message == nil {
					continue
				}
				g.P("if _val , ok := _opt[\"", oneofField.GoName, "\"].(map[string]interface{}); ok {")
				g.P("field := new(", oneofField.Message.GoIdent.GoName, ")")
				g.P("field.LoadMap(_val)")
				g.P("e.", oneof.GoName, " = &", oneofField.GoIdent, "{", oneofField.GoName, ":field}")
				g.P("}")
			}
			g.P("}")
		}
		g.P("}")
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
