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
	layout := "2006-01-02"
	switch t := value.(type) {
	case *DateFromString:
		*i = *t
	case DateFromString:
		*i = t
	case string:
		date, err := time.Parse(time.RFC3339, t)
		if err != nil {
			date, err = time.Parse(layout, t)
			if err != nil {
				return err
			}
		}
		*i = DateFromString(date)
	case time.Time:
		*i = DateFromString(t)
	case *string:
		if t != nil {
			date, err := time.Parse(time.RFC3339, *t)
			if err != nil {
				date, err = time.Parse(layout, *t)
				if err != nil {
					return err
				}
			}
			*i = DateFromString(date)
		}
	case *time.Time:
		if t != nil {
			*i = DateFromString(*t)
		}
	default:
		return errors.New(fmt.Sprint("failed to unmarshal IntToDate value:", value))
	}
	return nil
}

func (i DateFromString) Value() (driver.Value, error) {
	if time.Time(i).IsZero() {
		return nil, nil
	}
	return time.Time(i).Format(time.RFC3339), nil
}
