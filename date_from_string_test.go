package jsontostruct

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateFromString(t *testing.T) {
	// Test MarshalJSON
	date := DateFromString(time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	jsonStr := `{"field": "2023-01-07"}`
	_, err := json.Marshal(jsonStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test UnmarshalJSON
	var newDate DateFromString
	err = json.Unmarshal([]byte(`"2022-02-10"`), &newDate)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedDate := time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC)
	if time.Time(newDate) != expectedDate {
		t.Errorf("Unexpected date: got %v, want %v", newDate, expectedDate)
	}

	// Test Scan
	var scanDate DateFromString
	err = scanDate.Scan(time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if time.Time(scanDate) != time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Unexpected date: got %v, want %v", scanDate, time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	}

	// Test Value
	value, err := date.Value()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value.(time.Time) != date.Time() {
		t.Errorf("Unexpected value: got %v, want %v", value, date.Time())
	}
}
