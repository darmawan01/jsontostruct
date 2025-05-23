package jsontostruct

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Int64Encrypted int64

func (c Int64Encrypted) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Int64Encrypted) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if strings.HasPrefix(value, "0x02") {
		decoded, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
		if err != nil {
			return err
		}
		decrypedValue, err := DecryptByPassphrase(decoded)
		if err != nil {
			return err
		}

		i, err := strconv.Atoi(decrypedValue)
		if err != nil {
			return err
		}

		*c = Int64Encrypted(i)
		return nil
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		*c = Int64Encrypted(0)
	}

	*c = Int64Encrypted(i)

	return nil
}

func (c *Int64Encrypted) IsNil() bool {
	return c == nil
}

func (c *Int64Encrypted) String() string {
	if c.IsNil() {
		return "0"
	}

	return fmt.Sprintf("%d", *c)
}

func (c *Int64Encrypted) Scan(value interface{}) error {
	switch t := value.(type) {
	case *Int64Encrypted:
		*c = *t
	case Int64Encrypted:
		*c = t
	case int64:
		*c = Int64Encrypted(t)
	case *int64:
		*c = Int64Encrypted(*t)
	default:
		return errors.New(fmt.Sprint("failed to unmarshal Int64Encrypted value:", value))
	}

	return nil
}

func (c Int64Encrypted) Value() (driver.Value, error) {
	return int64(c), nil
}
