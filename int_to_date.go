package jsontostruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type IntToDate time.Time

func (i IntToDate) MarshalJSON() ([]byte, error) {
	val, err := strconv.Atoi(time.Time(i).Format("20060102"))
	if err != nil {
		return nil, err
	}
	return json.Marshal(val)

}

func (i *IntToDate) UnmarshalJSON(data []byte) error {
	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	dateString := fmt.Sprintf("%d", value)
	layout := "20060102"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return err
	}

	*i = IntToDate(date)
	return nil
}

func (i *IntToDate) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal IntToDate value:", value))
	}
	*i = IntToDate(t)

	return nil
}

func (c IntToDate) Value() (driver.Value, error) {

	return c, nil
}
