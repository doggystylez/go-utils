package json

import (
	"encoding/json"
	"os"
	"testing"
)

type testStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestLoadFile(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "*.json")
	defer os.Remove(tempFile.Name())
	ts := &testStruct{
		Field1: "test",
		Field2: 123,
	}
	json.NewEncoder(tempFile).Encode(ts)
	tempFile.Close()
	var result testStruct
	err := LoadFile(tempFile.Name(), &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Field1 != ts.Field1 || result.Field2 != ts.Field2 {
		t.Fatalf("expected %v, got %v", ts, result)
	}
}

func TestSaveFile(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "*.json")
	defer os.Remove(tempFile.Name())
	ts := &testStruct{
		Field1: "test",
		Field2: 123,
	}
	err := SaveFile(tempFile.Name(), ts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	var result testStruct
	file, _ := os.Open(tempFile.Name())
	json.NewDecoder(file).Decode(&result)
	file.Close()
	if result.Field1 != ts.Field1 || result.Field2 != ts.Field2 {
		t.Fatalf("expected %v, got %v", ts, result)
	}
}
