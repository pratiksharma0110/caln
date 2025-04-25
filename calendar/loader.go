package calendar

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseJson(path string) ([][]map[string]any, error) {

	username := os.Getenv("USER")
	absPath := "/home/" + username + "/Desktop/proj/caln/data/2082.json"

	jsonFile, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var data [][]map[string]any
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %v", err)
	}

	return data, nil
}
