package main

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type overrideJsonName struct {
	name string
	protoreflect.FieldDescriptor
}

func (d overrideJsonName) Name() protoreflect.Name {
	return protoreflect.Name(d.name)
}
