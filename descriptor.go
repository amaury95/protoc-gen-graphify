package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func updateMessageNames(messages ...*protogen.Message) {
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.Desc.HasJSONName() {
				jsonTag := field.Desc.JSONName()
				parts := strings.SplitN(jsonTag, ",", 2)
				if len(parts) == 0 {
					continue
				}
				field.Desc = overrideJsonName{name: parts[0], FieldDescriptor: field.Desc}
			}
		}
	}
}

type overrideJsonName struct {
	name string
	protoreflect.FieldDescriptor
}

func (d overrideJsonName) Name() protoreflect.Name {
	return protoreflect.Name(d.name)
}
