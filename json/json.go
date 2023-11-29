package json

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// LoadFile loads a json file into a new struct
func LoadFile[T any](filename string) (*T, error) {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	var s T
	err = json.NewDecoder(file).Decode(&s)
	if err != nil {
		return nil, err
	}
	_ = file.Close()
	return &s, nil
}

// MergeFile takes a json file and loads it into an existing struct
func MergeFile(filename string, s any) error {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(s)
	if err != nil {
		return err
	}
	_ = file.Close()
	return nil
}

// SaveFile saves a struct to a json file
func SaveFile(filename string, s any) error {
	file, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(s)
	_ = file.Close()
	return err
}

// Pretty returns a pretty-printed json string
func Pretty(s any) (string, error) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
