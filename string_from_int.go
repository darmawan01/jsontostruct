package jsontostruct

import (
	"encoding/json"
	"strconv"
)

type StringFromInt string

func (s StringFromInt) MarshalJSON() ([]byte, error) {

	val, err := strconv.Atoi(string(s))
	if err != nil {
		return nil, err
	}

	return json.Marshal(val)

}

func (s *StringFromInt) UnmarshalJSON(data []byte) error {
	if data != nil {
		value := string(data)

		*s = StringFromInt(value)
	}

	return nil
}
