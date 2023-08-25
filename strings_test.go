package jsontostruct

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var (
	key   = []byte("0123456789ABCDEF")
	value = "secret data"
)

func EncryptAESV2() (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(value), nil)

	// Combine the nonce and ciphertext into a single byte slice
	encryptedData := append(nonce, ciphertext...)

	// Encode the encrypted data in base64 for storage or transmission
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedValue, nil
}

func TestDecryptedStructOnUnmarshal(t *testing.T) {
	encryptedValue, err := EncryptAESV2()
	if err != nil {
		t.Fatal(err)
	}
	expectedDecryptedValue := value

	// Create a JSON string with the encrypted value
	jsonStr := `{"field1": "` + encryptedValue + `"}`

	// Unmarshal the JSON string into a struct
	var myStruct struct {
		Field1 String `json:"field1"`
	}
	err = json.Unmarshal([]byte(jsonStr), &myStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the decrypted value matches the expected value
	if strings.EqualFold(myStruct.Field1.String(), expectedDecryptedValue) {
		t.Errorf("Unexpected decrypted value: got %s, want %s", myStruct.Field1, expectedDecryptedValue)
	}
}
