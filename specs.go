package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func generateSpecs(g *protogen.GeneratedFile, _ *protogen.File, messages ...*protogen.Message) {
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}
		g.P("\n// Specs ...")
		g.P("func (*", message.GoIdent, ") Specs() []byte {")
		g.P("var buffer ", g.QualifiedGoIdent(bytesBuffer))
		g.P(bufferWrite("{")...)

		// fields
		g.P(bufferWrite(quote("fields"), ": {")...)
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}

			g.P(bufferWrite(quote(field.Desc.JSONName()), ": {")...)

			if field.Desc.HasPresence() {
				g.P(bufferWrite(quote("optional"), ": true,")...)
			}

			if field.Desc.IsList() {
				g.P(bufferWrite(quote("kind"), ":", quote("list"), ",")...)
			}

			if field.Desc.IsMap() {
				g.P(bufferWrite(quote("kind"), ":", quote("map"), ",")...)

				keyField := field.Message.Fields[0]
				g.P(bufferWrite(quote("key"), ": {")...)
				writeField(g, keyField)
				g.P(bufferWrite("},")...)

				valField := field.Message.Fields[1]
				g.P(bufferWrite(quote("value"), ": {")...)
				writeField(g, valField)
				g.P(bufferWrite("},")...)
			} else {
				writeField(g, field)
				g.P(bufferWrite(",")...)
			}

			g.P(bufferWrite(quote("name"), ":", quote(field.Desc.JSONName()))...)

			g.P(bufferWrite("},")...)

		}
		g.P(g.QualifiedGoIdent(trimTrailingComma), "(&buffer)")
		g.P(bufferWrite("},")...)

		// oneofs
		g.P(bufferWrite(quote("oneofs"), ": {")...)
		for _, field := range message.Oneofs {
			if field.Desc.IsSynthetic() {
				continue
			}
			g.P(bufferWrite(quote(field.GoName), ": {")...)
			for _, option := range field.Fields {
				g.P(bufferWrite(quote(option.Desc.JSONName()), ": ")...)
				g.P("_", option.GoName, " := new(", g.QualifiedGoIdent(option.Message.GoIdent), ")")
				g.P("buffer.Write(_", option.GoName, ".Specs())")
				g.P(bufferWrite(",")...)
			}
			g.P(g.QualifiedGoIdent(trimTrailingComma), "(&buffer)")
			g.P(bufferWrite("}")...)
		}
		g.P(bufferWrite("}")...)

		g.P(bufferWrite("}")...)
		g.P("return buffer.Bytes()")
		g.P("}")
	}
}

func writeField(g *protogen.GeneratedFile, field *protogen.Field) {
	if field.Desc.Kind() == protoreflect.MessageKind {
		fieldName := "_" + field.GoName
		g.P("var ", fieldName, " interface{} = new(", g.QualifiedGoIdent(field.Message.GoIdent), ")")
		g.P("if spec, ok := ", fieldName, ".(", g.QualifiedGoIdent(specs), "); ok {")
		g.P(bufferWrite(quote("value"), ":")...)
		g.P("buffer.Write(spec.Specs())")
		g.P("}")
		g.P(bufferWrite(",")...)
	}
	g.P(bufferWrite(quote("type"), ": ", quote(field.Desc.Kind().String()))...)
}

func bufferWrite(v ...interface{}) []interface{} {
	return join("buffer.WriteString(\"", v, "\")")
}

func quote(val string) string {
	return "\\\"" + val + "\\\""
}

var bytesBuffer protogen.GoIdent = protogen.GoIdent{
	GoName:       "Buffer",
	GoImportPath: "bytes",
}

var trimTrailingComma protogen.GoIdent = protogen.GoIdent{
	GoName:       "TrimTrailingComma",
	GoImportPath: "github.com/amaury95/protoc-gen-go-tag/utils",
}

var specs protogen.GoIdent = protogen.GoIdent{
	GoName:       "ISpecs",
	GoImportPath: "github.com/amaury95/protoc-gen-go-tag/utils",
}
