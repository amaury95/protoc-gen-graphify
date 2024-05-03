package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func generateObjects(g *protogen.GeneratedFile, file *protogen.File, messages ...*protogen.Message) {
	g.P(`
	/*
		Graphql object
	*/
	`)
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}

		g.P("\n/* QueryObject ... */")
		g.P("func (*", message.GoIdent, ") QueryObject() *", g.QualifiedGoIdent(Object), " {")
		g.P("return ", message.GoIdent, "_Object")
		g.P("}")

		g.P("var ", message.GoIdent, "_Object = ", g.QualifiedGoIdent(NewObject), "(", g.QualifiedGoIdent(ObjectConfig), "{")
		g.P("Name: \"", message.GoIdent, "\",")
		g.P("Fields: ", g.QualifiedGoIdent(Fields), "{")
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}
			g.P("\"", field.Desc.Name(), "\":&", g.QualifiedGoIdent(Field), "{")
			fieldType(g, field)
			g.P()
			g.P("},")
		}
		for _, field := range message.Oneofs {
			if field.Desc.IsSynthetic() {
				continue
			}

			g.P("\"", field.GoName, "\":&", g.QualifiedGoIdent(Field), "{")
			g.P("Type: ", g.QualifiedGoIdent(NewUnion), "(", g.QualifiedGoIdent(UnionConfig), "{")
			g.P("Name: \"", g.QualifiedGoIdent(field.GoIdent), "\",")
			g.P("Types: []*", g.QualifiedGoIdent(Object), "{")
			for _, option := range field.Fields {
				g.P("option_", g.QualifiedGoIdent(option.GoIdent), ",")
			}
			g.P("},")
			g.P("ResolveType: func(p ", g.QualifiedGoIdent(ResolveTypeParams), ") *", g.QualifiedGoIdent(Object), " {")
			g.P("switch   p.Value.(type) {")
			for _, option := range field.Fields {
				g.P("case *", g.QualifiedGoIdent(option.GoIdent), ":")
				g.P("return option_", g.QualifiedGoIdent(option.GoIdent))
			}
			g.P("default:")
			g.P("return nil")
			g.P("}")
			g.P("},")
			g.P("}),")
			g.P("},")
		}
		g.P("},")
		g.P("Description: \"\",")
		g.P("})")

		for _, oneof := range message.Oneofs {
			for _, option := range oneof.Fields {
				g.P()
				g.P("var option_", g.QualifiedGoIdent(option.GoIdent), " = ", g.QualifiedGoIdent(NewObject), "(", g.QualifiedGoIdent(ObjectConfig), "{")
				g.P("Name: \"", g.QualifiedGoIdent(option.GoIdent), "\",")
				g.P("Fields: ", g.QualifiedGoIdent(Fields), "{")
				g.P("\"", option.GoName, "\": &", g.QualifiedGoIdent(Field), "{")
				if option.Message != nil {
					g.P("Type: ", g.QualifiedGoIdent(option.Message.GoIdent), "_Object,")
				} else {
					fieldType(g, option)
				}
				g.P("},")
				g.P("},")
				g.P("})")
			}
		}
	}

	// usedEnums := map[string]bool{}
	// for _, field := range walkEnums(messages) {
	// 	// avoid repeated enums
	// 	if usedEnums[field.GoName] {
	// 		continue
	// 	}
	// 	usedEnums[field.GoName] = true

	// 	// generate enum field
	// 	g.P(`

	// 	`)
	// 	g.P("var ", field.GoName, "_Enum = ", g.QualifiedGoIdent(NewEnum), "(", g.QualifiedGoIdent(EnumConfig), "{")
	// 	g.P("Name: \"", field.GoName, "\",")
	// 	g.P("Values: ", g.QualifiedGoIdent(EnumValueConfigMap), "{")
	// 	for _, option := range field.Enum.Values {
	// 		g.P("\"", option.Desc.Name(), "\": &", g.QualifiedGoIdent(EnumValueConfig), "{ Value: ", g.QualifiedGoIdent(option.GoIdent), " },")
	// 	}
	// 	g.P("},")
	// 	g.P("})")
	// }

	for _, enum := range file.Enums {
		// generate enum field
		g.P(`
		
		`)
		g.P("var ", enum.Desc.Name(), "_Enum = ", g.QualifiedGoIdent(NewEnum), "(", g.QualifiedGoIdent(EnumConfig), "{")
		g.P("Name: \"", enum.Desc.Name(), "\",")
		g.P("Values: ", g.QualifiedGoIdent(EnumValueConfigMap), "{")
		for _, option := range enum.Values {
			g.P("\"", option.Desc.Name(), "\": &", g.QualifiedGoIdent(EnumValueConfig), "{ Value: ", g.QualifiedGoIdent(option.GoIdent), " },")
		}
		g.P("},")
		g.P("})")
	}
}

func walkEnums(messages []*protogen.Message) (result []*protogen.Field) {
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.Desc.Kind() == protoreflect.EnumKind {
				result = append(result, field)
			}
		}
		result = append(result, walkEnums(message.Messages)...)
	}
	return
}

func fieldType(g *protogen.GeneratedFile, field *protogen.Field) {
	fT := getFieldType(g, field)
	if fT != "" {
		if field.Desc.IsMap() {
			g.P("Type: ", g.QualifiedGoIdent(JSON), ",")
		} else if field.Desc.IsList() {
			g.P("Type: ", g.QualifiedGoIdent(NewList), "(", fT, "),")
		} else {
			g.P("Type: ", fT, ",")
		}
	}
}

func getFieldType(g *protogen.GeneratedFile, field *protogen.Field) string {
	if field.Desc.Name() == "_key" {
		return g.QualifiedGoIdent(ID)
	}
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		return g.QualifiedGoIdent(field.Message.GoIdent) + "_Object"
	case protoreflect.EnumKind:
		return g.QualifiedGoIdent(field.Enum.GoIdent) + "_Enum"
	case protoreflect.BoolKind:
		return g.QualifiedGoIdent(Boolean)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return g.QualifiedGoIdent(Int)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return g.QualifiedGoIdent(Int)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return g.QualifiedGoIdent(Int)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return g.QualifiedGoIdent(Int)
	case protoreflect.FloatKind:
		return g.QualifiedGoIdent(Float)
	case protoreflect.DoubleKind:
		return g.QualifiedGoIdent(Float)
	case protoreflect.StringKind:
		return g.QualifiedGoIdent(String)
	case protoreflect.BytesKind:
		return g.QualifiedGoIdent(Bytes)
	}
	return ""
}

var Fields protogen.GoIdent = protogen.GoIdent{
	GoName:       "Fields",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var Field protogen.GoIdent = protogen.GoIdent{
	GoName:       "Field",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var ID protogen.GoIdent = protogen.GoIdent{
	GoName:       "ID",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}
var String protogen.GoIdent = protogen.GoIdent{
	GoName:       "String",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var Int protogen.GoIdent = protogen.GoIdent{
	GoName:       "Int",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var Float protogen.GoIdent = protogen.GoIdent{
	GoName:       "Float",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var Boolean protogen.GoIdent = protogen.GoIdent{
	GoName:       "Boolean",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var DateTime protogen.GoIdent = protogen.GoIdent{
	GoName:       "DateTime",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var JSON protogen.GoIdent = protogen.GoIdent{
	GoName:       "JSON",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var Bytes protogen.GoIdent = protogen.GoIdent{
	GoName:       "Bytes",
	GoImportPath: protogen.GoImportPath("github.com/amaury95/protoc-gen-graphify/utils"),
}

var Object protogen.GoIdent = protogen.GoIdent{
	GoName:       "Object",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var NewObject protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewObject",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var ObjectConfig protogen.GoIdent = protogen.GoIdent{
	GoName:       "ObjectConfig",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var NewUnion protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewUnion",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var UnionConfig protogen.GoIdent = protogen.GoIdent{
	GoName:       "UnionConfig",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var NewScalar protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewScalar",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var ScalarConfig protogen.GoIdent = protogen.GoIdent{
	GoName:       "ScalarConfig",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var NewEnum protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewEnum",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var EnumConfig protogen.GoIdent = protogen.GoIdent{
	GoName:       "EnumConfig",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var EnumValueConfigMap protogen.GoIdent = protogen.GoIdent{
	GoName:       "EnumValueConfigMap",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var EnumValueConfig protogen.GoIdent = protogen.GoIdent{
	GoName:       "EnumValueConfig",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var NewList protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewList",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var ResolveTypeParams protogen.GoIdent = protogen.GoIdent{
	GoName:       "ResolveTypeParams",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}
