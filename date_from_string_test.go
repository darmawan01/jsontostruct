package jsontostruct

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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
	expectedDate := time.Date(2022, time.February, 10, 0, 0, 0, 0, time.UTC)
	if !expectedDate.Equal(time.Time(newDate)) {
		t.Errorf("Unexpected date: got %v, want %v", time.Time(newDate), expectedDate)
	}

	// Test Scan
	var scanDate DateFromString
	err = newDate.Scan(&scanDate)
	require.NoError(t, err)
	require.Equal(t, time.Time(scanDate), time.Time(newDate))

	// Test Value
	value, err := date.Value()
	require.NoError(t, err)

	require.Equal(t, value.(time.Time), time.Time(date))
}
