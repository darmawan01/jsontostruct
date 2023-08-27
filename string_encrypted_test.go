package jsontostruct

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"strings"
	"testing"
)

var (
	secretKey = []byte("secret key")
	secretVal = []byte("secret data")
)

// padPKCS7 pads the data with PKCS7 padding scheme to match the block size.
func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	paddedData := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return paddedData
}

// EncryptByPassphrase encrypts the plaintext using the passphrase and returns the ciphertext.
func encrypt() ([]byte, error) {
	plaintext := string(secretVal)
	key := passphraseToKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Prepend the magic number and authenticator to the encrypted data
	encryptedData := make([]byte, 20+len(plaintext))
	binary.LittleEndian.PutUint32(encryptedData[:4], magicNum)
	binary.LittleEndian.PutUint16(encryptedData[4:6], 0)

	plaintextBytes := []byte(plaintext)
	paddedData := padPKCS7(plaintextBytes, block.BlockSize())
	copy(encryptedData[20:], paddedData)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encryptedData[20:], paddedData)

	return encryptedData, nil
}

func TestEncryptedStructOnMarshal(t *testing.T) {
	SetPasshrase(string(secretKey))

	encryptedValue, err := encrypt()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(encryptedValue, "<============")
}

func TestDecryptedStructOnUnmarshal(t *testing.T) {
	SetPasshrase(string(secretKey))

	encryptedValue, err := encrypt()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(encryptedValue, "<============")
	expectedDecryptedValue := string(secretVal)

	// Create a JSON string with the encrypted value
	jsonStr := `{"field": "` + string(encryptedValue) + `"}`

	// Unmarshal the JSON string into a struct
	var myStruct struct {
		Field StringEncrypted `json:"field"`
	}
	err = json.Unmarshal([]byte(jsonStr), &myStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	log.Println(encryptedValue, expectedDecryptedValue, myStruct.Field)

	// Check if the decrypted value matches the expected value
	if !strings.EqualFold(myStruct.Field.String(), expectedDecryptedValue) {
		t.Errorf("Unexpected decrypted value: got %s, want %s", myStruct.Field, expectedDecryptedValue)
	}
}
