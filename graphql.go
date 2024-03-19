package main

import "google.golang.org/protobuf/compiler/protogen"

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
		g.P("func (*", message.GoIdent, ") Query(name string) *", g.QualifiedGoIdent(Object), " {")
		g.P("return ", g.QualifiedGoIdent(NewObject), "(", g.QualifiedGoIdent(ObjectConfig), "{")
		g.P("Name: name,")
		g.P("Fields: ", g.QualifiedGoIdent(Fields), "{")
		for _, field := range message.Fields {
			g.P("\"", field.Desc.Name(), "\":&", g.QualifiedGoIdent(Field), "{")
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
			g.P("Types: []*graphql.Object{")
			for _, option := range field.Fields {
				g.P("new(", g.QualifiedGoIdent(option.Message.GoIdent), ").Query(name+\"", option.GoName, "\"),")
			}
			g.P("},")
			g.P("ResolveType: func(p ", g.QualifiedGoIdent(ResolveTypeParams), ") *", g.QualifiedGoIdent(Object), " {")
			g.P("switch p.Value.(type) {")
			// 	g.P("case Dog:")
			// 		g.P("return dogType")
			// 	g.P("case Cat:")
			// 		g.P("return catType")
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
		g.P("}")
	}
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
