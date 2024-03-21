package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
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

// UnmarshalJSON ...
func (m *Map) UnmarshalJSON(b []byte) error {
	m.bytes = b
	return json.Unmarshal(b, &m.values)
}

// MarshalJSON ...
func (m *Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.values)
}

// GraphqlQuery ...
type GraphqlQuery interface {
	// QueryObject ...
	QueryObject() *graphql.Object
}

// GraphqlMutation ...
type GraphqlMutation interface {
	// MutationObject ...
	MutationObject() *graphql.Object
}

// JSON json type
var JSON = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "JSON",
		Description: "The `JSON` scalar type represents JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf)",
		Serialize: func(value interface{}) interface{} {
			return value
		},
		ParseValue: func(value interface{}) interface{} {
			return value
		},
		ParseLiteral: parseLiteral,
	},
)

func parseLiteral(astValue ast.Value) interface{} {
	kind := astValue.GetKind()

	switch kind {
	case kinds.StringValue:
		return astValue.GetValue()
	case kinds.BooleanValue:
		return astValue.GetValue()
	case kinds.IntValue:
		return astValue.GetValue()
	case kinds.FloatValue:
		return astValue.GetValue()
	case kinds.ObjectValue:
		obj := make(map[string]interface{})
		for _, v := range astValue.GetValue().([]*ast.ObjectField) {
			obj[v.Name.Value] = parseLiteral(v.Value)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteral(v))
		}
		return list
	default:
		return nil
	}
}
