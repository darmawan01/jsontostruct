package jsontostruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DateFromInt time.Time

func (i DateFromInt) MarshalJSON() ([]byte, error) {
	val, err := strconv.Atoi(time.Time(i).Format("20060102"))
	if err != nil {
		return nil, err
	}
	return json.Marshal(val)

}

func (i *DateFromInt) UnmarshalJSON(data []byte) error {
	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	dateString := strings.Split(fmt.Sprintf("%f", value), ".")[0]
	layout := "20060102"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return err
	}

	*i = DateFromInt(date)
	return nil
}

func (i *DateFromInt) Time() time.Time {
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
