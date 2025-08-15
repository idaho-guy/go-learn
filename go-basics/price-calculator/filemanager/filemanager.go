package filemanager

import (
	"bufio"
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
