package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func generateObject(g *protogen.GeneratedFile, _ *protogen.File, messages ...*protogen.Message) {
	g.P(`
	/*
		Graphql object
	*/
	`)
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}

		g.P("\n/* Query ... */")
		g.P("func (*", message.GoIdent, ") Query() *", g.QualifiedGoIdent(Object), " {")
		g.P("return ", message.GoIdent, "_Object")
		g.P("}")

		g.P("var ", message.GoIdent, "_Object = ", g.QualifiedGoIdent(NewObject), "(", g.QualifiedGoIdent(ObjectConfig), "{")
		g.P("Name: \"", message.GoIdent, "\",")
		g.P("Fields: ", g.QualifiedGoIdent(Fields), "{")
		for _, field := range message.Fields {

			g.P("\"", field.Desc.Name(), "\":&", g.QualifiedGoIdent(Field), "{")
			if field.Desc.IsList() {
				g.P("Type: ", g.QualifiedGoIdent(NewList), "(", getFieldType(g, field), "),")
			} else {
				g.P("Type: ", getFieldType(g, field), ",")
			}
			g.P()
			g.P("},")
		}
		for _, field := range message.Oneofs {
			if field.Desc.IsSynthetic() {
				continue
			}
			g.P("\"", field.GoName, "\":&", g.QualifiedGoIdent(Field), "{")
			// g.P("\"", field.GoName, "\": map[string]interface{}{")
			// for _, option := range field.Fields {
			// 	g.P("\"", option.GoName, "\": new(", g.QualifiedGoIdent(option.Message.GoIdent), ").Schema(),")
			// }
			// g.P("},")
			g.P("Type: ", g.QualifiedGoIdent(NewUnion), "(", g.QualifiedGoIdent(UnionConfig), "{")
			g.P("Types: []*", g.QualifiedGoIdent(Object), "{")
			for _, option := range field.Fields {
				g.P(g.QualifiedGoIdent(option.Message.GoIdent), "_Object,")
			}
			g.P("},")
			g.P("ResolveType: func(p ", g.QualifiedGoIdent(ResolveTypeParams), ") *", g.QualifiedGoIdent(Object), " {")
			g.P("switch   p.Value.(type) {")
			for _, option := range field.Fields {
				g.P("case ", g.QualifiedGoIdent(option.Message.GoIdent), ":")
				g.P("return ", g.QualifiedGoIdent(option.Message.GoIdent), "_Object")
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
	}
}

func getFieldType(g *protogen.GeneratedFile, field *protogen.Field) string {
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		// return g.QualifiedGoIdent()
	case protoreflect.BoolKind:
		return g.QualifiedGoIdent(Boolean)
	case protoreflect.EnumKind:
		// assignEnum(export, g, field, recipient, assign, identifier...)
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
		return g.QualifiedGoIdent(String)
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

var ID protogen.GoIdent = protogen.GoIdent{
	GoName:       "ID",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var DateTime protogen.GoIdent = protogen.GoIdent{
	GoName:       "DateTime",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
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

var NewList protogen.GoIdent = protogen.GoIdent{
	GoName:       "NewList",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}

var ResolveTypeParams protogen.GoIdent = protogen.GoIdent{
	GoName:       "ResolveTypeParams",
	GoImportPath: protogen.GoImportPath("github.com/graphql-go/graphql"),
}
