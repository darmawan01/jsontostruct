package jsontostruct

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDecryptedStructOnUnmarshal(t *testing.T) {
	encryptedValue := "0x02000000BF87541AB3E917E2B5AC21075FFD338CE45A59B994D53F05DECDE386EBC85BEB"

	SetPasshrase("Test123#$")

	expectedDecryptedValue := "060"

	// Create a JSON string with the encrypted value
	jsonStr := `{"field": "` + string(encryptedValue) + `"}`

	// Unmarshal the JSON string into a struct
	var myStruct struct {
		Field StringEncrypted `json:"field"`
	}
	err := json.Unmarshal([]byte(jsonStr), &myStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the decrypted value matches the expected value
	if !strings.EqualFold(*myStruct.Field.String(), expectedDecryptedValue) {
		t.Errorf("Unexpected decrypted value: got %s, want %s", myStruct.Field, expectedDecryptedValue)
	}
}

func TestFloat64OnUnmarshal(t *testing.T) {
	// Create a JSON string with the encrypted value
	jsonStr := `{"field": 0.45}`

	// Unmarshal the JSON string into a struct
	var myStruct struct {
		Field StringEncrypted `json:"field"`
	}
	err := json.Unmarshal([]byte(jsonStr), &myStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the decrypted value matches the expected value
	if !strings.EqualFold(*myStruct.Field.String(), "0.45") {
		t.Errorf("Unexpected decrypted value: got %s, want %s", myStruct.Field, "0.45")
	}
}
