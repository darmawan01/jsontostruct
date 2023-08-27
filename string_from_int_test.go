package jsontostruct

import "testing"

func TestStringFromInt(t *testing.T) {
	// Test MarshalJSON
	str := StringFromInt("123")
	expectedJSON := []byte(`123`)
	jsonData, err := str.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonData) != string(expectedJSON) {
		t.Errorf("Unexpected JSON data: got %s, want %s", jsonData, expectedJSON)
	}

	// Test UnmarshalJSON
	var newStr StringFromInt
	err = newStr.UnmarshalJSON([]byte(`456`))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(newStr) != "456" {
		t.Errorf("Unexpected string: got %v, want %v", newStr, "456")
	}
}
