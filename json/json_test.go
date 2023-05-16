package json_test

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type TestStruct struct {
		A string `json:"a"`
		B int    `json:"b"`
	}

	data := TestStruct{
		A: "1",
		B: 2,
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	var ts1 TestStruct
	err = json.Unmarshal(marshal, &ts1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("unmarshal: %+v", ts1)

	// ----
	// Unmarshal parses the JSON-encoded data and stores the result
	// in the value pointed to by v. If v is nil or not a pointer,
	// Unmarshal returns an InvalidUnmarshalError.
	var ts3 *TestStruct
	err = json.Unmarshal(marshal, ts3)
	if err != nil {
		t.Error(err) // json: Unmarshal(nil *json_test.TestStruct)
	}
	t.Logf("unmarshal: %+v", ts3)

	// ----
	var ts2 TestStruct
	err = json.Unmarshal(marshal, ts2)
	if err != nil {
		t.Error(err) // json: Unmarshal(non-pointer json_test.TestStruct)
	}
	t.Logf("unmarshal: %+v", ts2)
}
