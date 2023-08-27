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

func (i *DateFromInt) Time() time.Time {
	if i == nil {
		return time.Time{}
	}

	return time.Time(*i)
}

func (i *DateFromInt) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal IntToDate value:", value))
	}

	*i = DateFromInt(t)
	return nil
}

func (i *DateFromInt) Value() (driver.Value, error) {
	return i.Time(), nil
}
