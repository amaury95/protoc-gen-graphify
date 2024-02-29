package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func updateMessageNames(messages ...*protogen.Message) {
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.Desc.HasJSONName() {
				field.Desc = overrideJsonName{name: field.Desc.JSONName(), FieldDescriptor: field.Desc}
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
