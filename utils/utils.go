package utils

import (
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

// MapFromBytes ...
func MapFromBytes(data []byte) (res map[string]interface{}, err error) {
	err = json.Unmarshal(data, &res)
	return
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
		ParseLiteral: parseLiteralJSON,
	},
)

func parseLiteralJSON(astValue ast.Value) interface{} {
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
			obj[v.Name.Value] = parseLiteralJSON(v.Value)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteralJSON(v))
		}
		return list
	default:
		return nil
	}
}

// Custom scalar type for []byte to base64 string and vice versa
var Bytes = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Bytes",
	Description: "Bytes scalar to encode []byte to base64 string and decode base64 string to []byte",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case []byte:
			return base64.StdEncoding.EncodeToString(value)
		case string:
			return value
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			return DecodeBytes(value)
		case []byte:
			return value
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})
