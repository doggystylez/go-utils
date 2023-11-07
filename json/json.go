package json

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// LoadFile loads a json file into a (pointer to a) struct
func LoadFile(filename string, s any) error {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(s)
	_ = file.Close()
	return err
}

// SaveFile saves a (pointer to a) struct into a json file
func SavFile(filename string, s any) error {
	file, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(s)
	_ = file.Close()
	return err
}
