package main

import "google.golang.org/protobuf/compiler/protogen"

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
		var fields [][]interface{}
		for _, field := range message.Fields {
			var rep []interface{}
			if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
				continue
			}

			rep = append(rep, quote(field.Desc.JSONName()), ": {")
			rep = append(rep, quote("name"), ":", quote(field.Desc.JSONName()), ",")
			rep = append(rep, quote("type"), ":", quote(field.Desc.Kind().String()))
			rep = append(rep, "}")

			fields = append(fields, rep)
		}
		for _, field := range joinLines(comma, fields...) {
			g.P(bufferWrite(field...)...)
		}
		g.P(bufferWrite("},")...)

		// oneofs
		g.P(bufferWrite(quote("oneofs"), ": {")...)
		var oneofs [][]interface{}
		for _, field := range message.Oneofs {
			var rep []interface{}
			if field.Desc.IsSynthetic() {
				continue
			}
			rep = append(rep, bufferWrite(quote(field.GoName), ": {")...)
			rep = append(rep, bufferWrite("}")...)
			oneofs = append(oneofs, rep)
		}
		for _, oneof := range joinLines(comma, oneofs...) {
			g.P(oneof...)
		}
		g.P(bufferWrite("}")...)

		g.P(bufferWrite("}")...)
		g.P("return buffer.Bytes()")
		g.P("}")
	}
}

func joinLines(sep []interface{}, vals ...[]interface{}) (result [][]interface{}) {
	for i, val := range vals {
		result = append(result, val)
		if i < len(vals)-1 {
			result = append(result, sep)
		}
	}
	return
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

var comma = []interface{}{","}
