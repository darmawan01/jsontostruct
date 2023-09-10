package jsontostruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DateFromString time.Time

func (i DateFromString) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(i).Format("2006-01-02")), nil
}

func (i *DateFromString) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	layout := "2006-01-02"
	date, err := time.Parse(layout, value)
	if err != nil {
		return err
	}

	*i = DateFromString(date)
	return nil
}

func (i *DateFromString) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal IntToDate value:", value))
	}

	*i = DateFromString(t)
	return nil
}

func (i DateFromString) Value() (driver.Value, error) {
	if time.Time(i).IsZero() {
		return "0000-00-00 00:00:00", nil
	}
	return time.Time(i), nil
}
