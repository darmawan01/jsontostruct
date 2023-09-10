package jsontostruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DateTimeFromDateTime time.Time

const layoutDtFromDt = "2006-01-02 15:04:05.0"

func (i DateTimeFromDateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(layoutDtFromDt)+2)
	b = append(b, '"')
	b = append(b, []byte(time.Time(i).Format(layoutDtFromDt))...)
	b = append(b, '"')
	return b, nil
}

func (i *DateTimeFromDateTime) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	date, err := time.Parse(layoutDtFromDt, value)
	if err != nil {
		return err
	}

	*i = DateTimeFromDateTime(date)
	return nil
}

func (i *DateTimeFromDateTime) Scan(value interface{}) error {
	t, ok := value.(*DateTimeFromDateTime)
	if !ok {
		return errors.New(fmt.Sprint("failed to scan DateTimeFromDateTime value:", value))
	}

	*i = DateTimeFromDateTime(*t)
	return nil
}

func (i DateTimeFromDateTime) Value() (driver.Value, error) {
	if time.Time(i).IsZero() {
		return "0000-00-00 00:00:00", nil
	}
	return time.Time(i), nil
}
