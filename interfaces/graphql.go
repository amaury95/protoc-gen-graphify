package interfaces

import "github.com/graphql-go/graphql"

// GraphqlObject ...
type GraphqlObject interface {
	// Object ...
	Object() *graphql.Object
}

// GraphqlArgument ...
type GraphqlArgument interface {
	// Argument ...
	Argument() graphql.FieldConfigArgument
}

// GraphqlOutput ...
type GraphqlOutput interface {
	// Output ...
	Output() graphql.Output
}