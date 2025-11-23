package main

import (
	"encoding/json"
)

// JsonSerializer provides static-style JSON serialization helpers.
type JsonSerializer struct{}

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
func Deserialize[T any](jsonText string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(jsonText), &result)
	return result, err
}

// Serialize serializes any Go value into JSON text.
func Serialize[T any](value T) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/* Usage example:
// Example type
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    serializer := JsonSerializer{}

    // Example JSON array
    jsonArray := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`

    // Deserialize list
    people, err := serializer.DeserializeList[Person](jsonArray)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("People list:", people)

    // Example JSON object
    jsonObject := `{"name":"Charlie","age":40}`

    // Deserialize single
    person, err := serializer.DeserializeSingle[Person](jsonObject)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Single person:", person)

    // Serialize back to JSON
    jsonText, err := serializer.Serialize(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Serialized JSON:", jsonText)
}
*/
