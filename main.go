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
		g.P("// LoadMap loads map values into given struct.")
		g.P("func (e *", message.GoIdent, ") LoadMap(m map[string]interface{}) {")
		for _, field := range message.Fields {
			if field.Oneof != nil {
				continue
			}
			goType, _ := fieldGoType(g, f, field)
			g.P("e."+field.GoName+" = m[\"", field.Desc.Name(), "\"]", ".(", goType, ")")
		}
		for _, oneof := range message.Oneofs {
			if oneof.Desc.IsSynthetic() {
				continue
			}
			g.P("if _opt, ok := m[\"" + oneof.GoName + "\"].(map[string]interface{}); ok {")
			for _, oneofField := range oneof.Fields {
				g.P("if _val , ok := _opt[\"", oneofField.GoName, "\"].(map[string]interface{}); ok {")
				g.P("field := new(", oneofField.Oneof.GoIdent, ")")
				g.P("field.LoadMap(_val)")
				g.P("//", oneofField)
				g.P("e.Type = &", oneofField.GoIdent, "{", oneofField.GoName, ":field}")
				g.P("}")
				/*
						if value , ok := option["Admin"].(map[string]interface{});ok {
						val:= new(Person_Admin)
						val.LoadMap(value)
						e.Type = &Person_Admin_{Admin: val}
					}
				*/

			}
			g.P("}")
		}
		g.P("}")
	}
}

// fieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGoType(g *protogen.GeneratedFile, f *protogen.File, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := fieldGoType(g, f, field.Message.Fields[0])
		valType, _ := fieldGoType(g, f, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}
