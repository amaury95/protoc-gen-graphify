package main

import (
	"reflect"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func exposeMapBuilders(g *protogen.GeneratedFile, _ *protogen.File, messages ...*protogen.Message) {
	g.P(`
	/*
		Graphify unmarshaler
	*/
	`)
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}

		g.P("\n/* UnmarshalJSON ...*/")
		g.P("func (o *", message.GoIdent, ") UnmarshalJSON(b []byte) error {")
		g.P("if values, err := ", g.QualifiedGoIdent(toMap), "(b);err != nil {return err} else {o.UnmarshalMap(values)}")
		g.P("return nil")
		g.P("}")

		g.P("\n/* UnmarshalMap populates struct fields from a map, handling decoding for special fields. */")
		g.P("func (o *", message.GoIdent, ") UnmarshalMap(values map[string]interface{}) {")
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
		if oneofField.Message != nil {
			g.P("if val , ok := _opt[\"", oneofField.GoName, "\"].(map[string]interface{}); ok {")
			g.P("field := new(", g.QualifiedGoIdent(oneofField.Message.GoIdent), ")")
			g.P("field.UnmarshalMap(val)")
			g.P(recipient, assign, "&", oneofField.GoIdent, "{", oneofField.GoName, ":field}")
			g.P("}")
		} else {
			g.P("if val, ok := _opt[\"", oneofField.GoName, "\"].(interface{}); ok {")
			fetchAndExportField(g, oneofField, recipient, assign, "val")
			g.P("}")
		}

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
		g.P(g.QualifiedGoIdent(makeSlice), "(&", recipient, ", len(list))")
		g.P("for index, item := range list {")
		fetchField(g, ignoreType(field), recipient+"[index]", "=", "item")
		g.P("}")
		g.P("}")
		return
	case field.Desc.IsMap():
		keyField := field.Message.Fields[0]
		valField := field.Message.Fields[1]
		g.P(join("if values, ok := ", identifier, ".(map[string]interface{}); ok {")...)
		g.P(g.QualifiedGoIdent(makeMap), "(&", recipient, ")")
		g.P("for key, value := range values {")
		if keyField.Desc.Kind() == protoreflect.StringKind {
			fetchField(g, valField, recipient+"[key]", "=", "value")
		} else {
			g.P("var tmp interface{} = ", g.QualifiedGoIdent(parseFloat), "(key)")
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
		parseBytes(export, g, field, recipient, assign, identifier...)
	}
}

func assignMessage(_ bool, g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	g.P(join("if val , ok := ", identifier, ".(map[string]interface{}); ok {")...)
	g.P("field := new(", g.QualifiedGoIdent(field.Message.GoIdent), ")")
	g.P("field.UnmarshalMap(val)")
	g.P(recipient, assign, "field")
	g.P("}")
}

func assignEnum(_ bool, g *protogen.GeneratedFile, field *protogen.Field, recipient, assign string, identifier ...interface{}) {
	parseField(true, g, field, "float64", "int32", field.Desc.JSONName(), " = ", identifier...)
	g.P(recipient, assign, g.QualifiedGoIdent(field.Enum.GoIdent), "(", field.Desc.JSONName(), ")")
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
func parseBytes(export bool, g *protogen.GeneratedFile, _ *protogen.Field, recipient, assign string, identifier ...interface{}) {
	if export {
		g.P("var ", recipient, " []byte")
	}
	g.P(join("if val, ok := ", identifier, ".(string); ok {")...)
	g.P(recipient, assign, g.QualifiedGoIdent(decodeBytes), "(val)")
	g.P("}")
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
	return
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

var decodeBytes protogen.GoIdent = protogen.GoIdent{
	GoName:       "DecodeBytes",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var makeMap protogen.GoIdent = protogen.GoIdent{
	GoName:       "MakeMap",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var makeSlice protogen.GoIdent = protogen.GoIdent{
	GoName:       "MakeSlice",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var parseFloat protogen.GoIdent = protogen.GoIdent{
	GoName:       "ParseFloat",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var toMap protogen.GoIdent = protogen.GoIdent{
	GoName:       "MapFromBytes",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}
