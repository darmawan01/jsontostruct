package jsontostruct

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type DateFromInt time.Time

func (i DateFromInt) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(i).Format("20060102")), nil
}

func (i *DateFromInt) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	value, err := strconv.Atoi(dataStr)
	if err != nil {
		return err
	}

	layout := "20060102"
	date, err := time.Parse(layout, fmt.Sprintf("%d", value))
	if err != nil {
		return err
	}

	*i = DateFromInt(date)
	return nil
}

func (i *DateFromInt) Scan(value interface{}) error {
	layout := "20060102"
	switch t := value.(type) {
	case string:
		date, err := time.Parse(time.RFC3339, t)
		if err != nil {
			date, err = time.Parse(layout, t)
			if err != nil {
				return err
			}
		}
		*i = DateFromInt(date)
	case time.Time:
		*i = DateFromInt(t)
	case *string:
		if t != nil {
			date, err := time.Parse(time.RFC3339, *t)
			if err != nil {
				date, err = time.Parse(layout, *t)
				if err != nil {
					return err
				}
			}
			*i = DateFromInt(date)
		}
	case *time.Time:
		if t != nil {
			*i = DateFromInt(*t)
		}
	default:
		return errors.New(fmt.Sprint("failed to unmarshal IntToDate value:", value))

	}

	return nil
}

func (i DateFromInt) Value() (driver.Value, error) {
	if time.Time(i).IsZero() {
		return nil, nil
	}
	return time.Time(i).Format(time.RFC3339), nil
}
