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

// func (c StringEncrypted) MarshalJSON() ([]byte, error) {
// 	return []byte(*c.String()), nil
// }

func (c *StringEncrypted) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		var value float64
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}

		*c = StringEncrypted(fmt.Sprintf("%v", value))
		return nil
	}

	if strings.HasPrefix(value, "0x02") {
		decoded, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
		if err != nil {
			return err
		}
		value, err = DecryptByPassphrase(decoded)
		if err != nil {
			return err
		}

	}

	var newVal string
	for _, v := range value {
		if v == 160 { // for &nsbp
			newVal += " "
			continue
		}
		newVal += string(v)
	}
	newVal = strings.ReplaceAll(newVal, "'", `"`)
	*c = StringEncrypted(newVal)

	return nil
}

func (c *StringEncrypted) IsNil() bool {
	return c == nil
}

func (c *StringEncrypted) String() *string {
	if c.IsNil() {
		return nil
	}

	val := string(*c)
	return &val
}

func (c *StringEncrypted) Scan(value interface{}) error {
	switch t := value.(type) {
	case string:
		*c = StringEncrypted(t)
	case []byte:
		*c = StringEncrypted(string(t))
	case *string:
		*c = StringEncrypted(*t)
	case *[]byte:
		*c = StringEncrypted(string(*t))
	default:
		return errors.New(fmt.Sprint("failed to unmarshal String value:", value))
	}

	return nil
}

func (c StringEncrypted) Value() (driver.Value, error) {
	return c.String(), nil
}
