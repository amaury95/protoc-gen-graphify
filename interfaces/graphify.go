package interfaces

// Unmarshaler ...
type Unmarshaler interface {
	UnmarshalMap(m map[string]interface{})
}

// Message ...
type Message interface {
	Schema() map[string]interface{}
}
