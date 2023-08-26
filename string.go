package jsontostruct

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

func (c *String) IsNil() bool {
	return c == nil
}

func (c *String) String() string {
	if c.IsNil() {
		return ""
	}

	return string(*c)
}

func (c *String) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal String value:", value))
	}

	*c = String(string(bytes))

	return nil
}

func (c String) Value() (driver.Value, error) {
	if len(c) == 0 {
		return nil, nil
	}

	return c.String(), nil
}
