package jsontostruct

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestDateFromInt(t *testing.T) {
	// Test MarshalJSON
	date := DateFromInt(time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	expectedJSON := []byte(`20230107`)
	jsonData, err := json.Marshal(date)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonData) != string(expectedJSON) {
		t.Errorf("Unexpected JSON data: got %s, want %s", jsonData, expectedJSON)
	}

	// Test UnmarshalJSON
	var newDate DateFromInt
	err = json.Unmarshal([]byte(`20230107`), &newDate)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if time.Time(newDate) != time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Unexpected date: got %v, want %v", newDate, time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	}

	// Test Scan
	var scanDate DateFromInt
	err = scanDate.Scan(time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if time.Time(scanDate) != time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Unexpected date: got %v, want %v", scanDate, time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC))
	}

	// Test Value
	_, err = date.Value()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	type test struct {
		Date DateFromInt `json:"date"`
	}

	var tests test
	jsonStr := `{"date": 20230107}`
	err = json.Unmarshal([]byte(jsonStr), &tests)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	log.Println(tests.Date.Value())

}
