package jsontostruct

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntEncrypted(t *testing.T) {
	encryptedValue := "0x02000000E16B019A1A1299744DEA2DD7B7A84A1A5569D8AAF5D5D09C21B7F871F4312D3589A97635FFDB997D297C39C233C14E95"
	expected := 363301016

	SetPasshrase("Test123#$")

	// decoded, err := hex.DecodeString(strings.TrimPrefix(encryptedValue, "0x"))
	// require.NoError(t, err)
	// s, err := DecryptByPassphrase(decoded)
	// require.Equal(t, nil, err)
	// fmt.Printf("s: %v\n", s)

	// Create a JSON string with the encrypted value
	jsonStr := `{"field": "` + string(encryptedValue) + `"}`

	// Unmarshal the JSON string into a struct
	var myStruct struct {
		Field Int64Encrypted `json:"field"`
	}

	err := json.Unmarshal([]byte(jsonStr), &myStruct)
	require.NoError(t, err)

	// Check if the decrypted value matches the expected value
	require.Equal(t, expected, int(myStruct.Field))

	v, err := myStruct.Field.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%d", expected), string(v))
}
