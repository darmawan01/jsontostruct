package jsontostruct

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StringEncrypted string

func (c StringEncrypted) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *StringEncrypted) UnmarshalJSON(data []byte) error {
	value := string(data)

	if strings.HasPrefix(value, "0x") {
		decoded, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
		if err != nil {
			return err
		}
		decrypedValue, err := DecryptByPassphrase(decoded)
		if err != nil {
			return err
		}

		*c = StringEncrypted(decrypedValue)
	}

	return nil
}

func (c *StringEncrypted) IsNil() bool {
	return c == nil
}

func (c *StringEncrypted) String() string {
	if c.IsNil() {
		return ""
	}

	return string(*c)
}

func (c *StringEncrypted) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal String value:", value))
	}

	*c = StringEncrypted(string(bytes))

	return nil
}

func (c *StringEncrypted) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}

	return c.String(), nil
}
