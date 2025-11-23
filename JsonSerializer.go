package main

import (
	"encoding/json"
	"fmt"
)

// JSONBox holds any generic value of type T.
type JSONBox[T any] struct {
	Value T
}

// Serialize method for JSONBox,
// internally delegates to the static Serialize function.
func (b JSONBox[T]) Serialize() ([]byte, error) {
	return Serialize(b.Value)
}

// Deserialize method for JSONBox,
// internally delegates to the static Deserialize function.
func (b *JSONBox[T]) Deserialize(data []byte) error {
	v, err := Deserialize[T](data)
	if err != nil {
		return err
	}
	b.Value = v
	return nil
}

// Serialize behaves like a static method for any type T.
func Serialize[T any](v T) ([]byte, error) {
	return json.Marshal(v)
}

// Deserialize behaves like a static method for any type T.
func Deserialize[T any](data []byte) (T, error) {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return v, fmt.Errorf("deserialize error: %w", err)
	}
	return v, nil
}

// DeserializeList deserializes JSON text into a slice of type T.
func DeserializeList[T any](jsonText string) ([]T, error) {
	var result []T
	err := json.Unmarshal([]byte(jsonText), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeserializeSingle deserializes JSON text into a single object of type T.
func DeserializeFromString[T any](jsonText string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(jsonText), &result)
	return result, err
}

// Serialize serializes any Go value into JSON text.
func SerializeToString[T any](value T) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
