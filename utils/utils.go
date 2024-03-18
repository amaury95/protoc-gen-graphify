package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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

// MapFromBytes ...
func MapFromBytes(data []byte) (res map[string]interface{}, err error) {
	err = json.Unmarshal(data, &res)
	return
}

// Unmarshaler ...
type Unmarshaler interface {
	UnmarshalMap(m map[string]interface{})
}

// Message ...
type Message interface {
	Schema() map[string]interface{}
}

// Map ...
type Map struct {
	values map[string]interface{}
	bytes  []byte
}

// Values ...
func (m *Map) Values() map[string]interface{} { return m.values }

// Bytes ...
func (m *Map) Bytes() []byte { return m.bytes }

func (m *Map) UnmarshalJSON(b []byte) error {
	m.bytes = b
	return json.Unmarshal(b, &m.values)
}

func (m *Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.values)
}
