package utils

import (
	"encoding/json"
	"errors"
)

// EncodeToJSON converts a given data into its JSON representation.
func EncodeToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// DecodeFromJSON decodes a JSON representation into the given data structure.
func DecodeFromJSON(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

// IsValidJSON checks if the given data is a valid JSON.
func IsValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}