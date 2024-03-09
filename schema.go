package main

import (
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func generateSchema(g *protogen.GeneratedFile, _ *protogen.File, messages ...*protogen.Message) {
	g.P(`
	/*
		Graphify schema module
	*/
	`)
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}
		g.P("\n// Specs ...")
		g.P("func (*", message.GoIdent, ") Schema() map[string]interface{} {")
		g.P("return map[string]interface{} {")

		g.P("\"fields\": map[string] interface{} {")

		// fields
		// g.P(bufferWrite(quote("fields"), ": {")...)
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}

			g.P("\"", field.Desc.Name(), "\": map[string]interface{}{")

			g.P("\"name\": \"", field.Desc.Name(), "\",")

			if field.Desc.HasPresence() {
				g.P("\"optional\": true,")
			}

			if field.Desc.IsList() {
				g.P("\"kind\": \"list\",")
			}

			if field.Desc.IsMap() {
				g.P("\"kind\": \"map\",")

				keyField := field.Message.Fields[0]
				g.P("\"key\": map[string]interface{}{")
				writeField(g, keyField)
				g.P("},")

				valField := field.Message.Fields[1]
				g.P("\"value\": map[string]interface{}{")
				writeField(g, valField)
				g.P("},")
			} else {
				writeField(g, field)
			}

			g.P("},")
		}
		g.P("},")

		// oneofs
		g.P("\"oneofs\": map[string] interface{} {")

		for _, field := range message.Oneofs {
			if field.Desc.IsSynthetic() {
				continue
			}

			g.P("\"", field.GoName, "\": map[string]interface{}{")
			for _, option := range field.Fields {
				g.P("\"", option.GoName, "\": new(", g.QualifiedGoIdent(option.Message.GoIdent), ").Schema(),")
			}
			g.P("},")
		}
		g.P("},")

		g.P("}")
		g.P("}")
	}
}

func writeField(g *protogen.GeneratedFile, field *protogen.Field) {
	if field.Desc.Kind() == protoreflect.EnumKind {
		g.P(bufferWrite(quote("options"), ": {")...)
		for index, option := range field.Enum.Values {
			g.P(bufferWrite(quote(strconv.Itoa(index)), ":", quote(string(option.Desc.Name())), ",")...)
		}
		g.P(g.QualifiedGoIdent(trimTrailingComma), "(&buffer)")
		g.P(bufferWrite("},")...)
	}
	if field.Desc.Kind() == protoreflect.MessageKind {
		fieldName := "_" + field.GoName
		g.P("var ", fieldName, " interface{} = new(", g.QualifiedGoIdent(field.Message.GoIdent), ")")
		g.P("if spec, ok := ", fieldName, ".(", g.QualifiedGoIdent(schema), "); ok {")
		g.P(bufferWrite(quote("schema"), ":")...)
		g.P("buffer.Write(spec.Schema())")
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
	GoImportPath: "github.com/amaury95/protoc-gen-graphify/utils",
}

var schema protogen.GoIdent = protogen.GoIdent{
	GoName:       "ISchema",
	GoImportPath: "github.com/amaury95/protoc-gen-graphify/utils",
}
