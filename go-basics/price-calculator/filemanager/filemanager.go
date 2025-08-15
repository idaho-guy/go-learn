package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read lines in file")
	}
	return lines, nil
}

func WritonJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to covert data to JSON")
	}
	file.Close()
	return nil
}
