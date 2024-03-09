package utils

import (
	"bytes"
	"encoding/base64"
	"reflect"
	"strconv"
)

// ParseFloat ...
func ParseFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

// MakeSlice ...
func MakeSlice(ptr interface{}, size int) {
	ptrVal := reflect.ValueOf(ptr)
	if ptrVal.Kind() == reflect.Ptr && ptrVal.Elem().Kind() == reflect.Slice {
		_n := reflect.MakeSlice(ptrVal.Elem().Type(), size, size)
		ptrVal.Elem().Set(_n)
	}
}

// MakeMap ...
func MakeMap(ptr interface{}) {
	mapVal := reflect.ValueOf(ptr)
	if mapVal.Kind() == reflect.Ptr && mapVal.Elem().Kind() == reflect.Map {
		newMap := reflect.MakeMap(mapVal.Elem().Type())
		mapVal.Elem().Set(newMap)
	}
}

// DecodeBytes ...
func DecodeBytes(input string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err == nil {
		return decoded
	}
	return []byte(input)
}

// TrimTrailingComma ...
func TrimTrailingComma(bb *bytes.Buffer) {
	if bb.Len() > 0 && bb.Bytes()[bb.Len()-1] == ',' {
		// Remove the last byte
		bb.Truncate(bb.Len() - 1)
	}
}

// IMapLoader ...
type IMapLoader interface {
	LoadMap(m map[string]interface{})
}

// ISchema ...
type ISchema interface {
	Schema() map[string]interface{}
}
