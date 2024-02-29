package main

import "google.golang.org/protobuf/compiler/protogen"

func generateSpecs(g *protogen.GeneratedFile, _ *protogen.File, messages ...*protogen.Message) {
	for _, message := range messages {
		if message.Desc.IsMapEntry() {
			continue
		}
		g.P()
		g.P("// Specs ...")
		g.P("func (*", message.GoIdent, ") Specs() []byte {")
		g.P("var buffer ", g.QualifiedGoIdent(bytesBuffer))
		for _, field := range message.Fields {
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}
			P(g, quote("name"), ":", quote(string(field.Desc.Name())))
		}
		P(g, "}")
		for _, field := range message.Oneofs {
			if field.Desc.IsSynthetic() {
				continue
			}
		}
		P(g, "}")
	}
	g.P("return buffer.Bytes()")
	g.P("}")
}

var bytesBuffer protogen.GoIdent = protogen.GoIdent{
	GoName:       "Buffer",
	GoImportPath: "bytes",
}

func P(g *protogen.GeneratedFile, v ...interface{}) {
	g.P(join("buffer.WriteString(\"", v, "\")")...)
}

func quote(val string) string {
	return "\\\"" + val + "\\\""
}
