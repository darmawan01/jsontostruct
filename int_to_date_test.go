package jsontostruct

import (
	"encoding/json"
	"testing"
	"time"
)

func TestIntToDate(t *testing.T) {
	// Test MarshalJSON
	date := IntToDate(time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	expectedJSON := []byte(`20230107`)
	jsonData, err := json.Marshal(date)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonData) != string(expectedJSON) {
		t.Errorf("Unexpected JSON data: got %s, want %s", jsonData, expectedJSON)
	}

	// Test UnmarshalJSON
	var newDate IntToDate
	err = json.Unmarshal([]byte(`20230107`), &newDate)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if time.Time(newDate) != time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Unexpected date: got %v, want %v", newDate, time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	}

	// Test Scan
	var scanDate IntToDate
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
	if value != date {
		t.Errorf("Unexpected value: got %v, want %v", value, date)
	}
}
