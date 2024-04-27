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
		g.P("\n/* Schema ... */")
		g.P("func (*", message.GoIdent, ") Schema() map[string]interface{} {")
		g.P("return map[string]interface{} {")
		g.P("\"name\": \"", message.GoIdent, "\",")
		g.P("\"fields\": []interface{} {")

		// fields
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}

			g.P("map[string]interface{}{")

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
				g.P("\"", option.GoName, "\": map[string]interface{} {")
				writeField(g, option)
				g.P("},")
			}
			g.P("},")
		}
		g.P("},")

		g.P("}")
		g.P("}")
	}
}

func writeField(g *protogen.GeneratedFile, field *protogen.Field) {
	g.P("\"type\": \"", field.Desc.Kind().String(), "\",")
	if field.Desc.Kind() == protoreflect.EnumKind {
		g.P("\"options\": map[string]interface{}{")
		for index, option := range field.Enum.Values {
			g.P("\"", strconv.Itoa(index), "\": \"", option.Desc.Name(), "\",")
		}
		g.P("},")
	}
	if field.Desc.Kind() == protoreflect.MessageKind {
		g.P("\"schema\": new(", g.QualifiedGoIdent(field.Message.GoIdent), ").Schema(),")
	}
}
