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
	if value == "" {
		return nil
	}
	date, err := time.Parse(layoutDtFromDt, value)
	if err != nil {
		return err
	}

	*i = DateTimeFromDateTime(date)
	return nil
}

func (i *DateTimeFromDateTime) Scan(value interface{}) error {
	switch t := value.(type) {
	case *DateTimeFromDateTime:
		*i = *t
	case DateTimeFromDateTime:
		*i = t
	case string:
		date, err := time.Parse(time.RFC3339, t)
		if err != nil {
			return err
		}

		*i = DateTimeFromDateTime(date)

	case time.Time:
		*i = DateTimeFromDateTime(t)

	case *string:
		if t != nil {
			date, err := time.Parse(time.RFC3339, *t)
			if err != nil {
				return err
			}
			*i = DateTimeFromDateTime(date)
		}
	case *time.Time:
		if t != nil {
			*i = DateTimeFromDateTime(*t)
		}

	default:
		return errors.New(fmt.Sprint("failed to unmarshal DateTimeFromDateTime value:", value))
	}

	return nil
}

func (i DateTimeFromDateTime) Value() (driver.Value, error) {
	if time.Time(i).IsZero() {
		return nil, nil
	}
	return time.Time(i), nil
}
