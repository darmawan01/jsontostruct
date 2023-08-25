package jsontostruct

import (
	"encoding/hex"
	"encoding/json"
	"strings"
)

type String string

func (c String) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *String) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if strings.HasPrefix(value, "0x") {
		decoded, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
		if err != nil {
			return err
		}
		decrypedValue, err := DecryptByPassphrase(decoded)
		if err != nil {
			return err
		}

		*c = String(decrypedValue)
	}

	return nil
}

func (c String) String() string {
	return string(c)
}
