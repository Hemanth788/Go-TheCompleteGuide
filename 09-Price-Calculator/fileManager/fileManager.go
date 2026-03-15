package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLinesFromFile(path string) ([]string, error) {
		file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Could not open the file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read the file content")
	}

	file.Close()
	return lines, nil
}

func WriteJSONToFile(data any, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create a file")
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	
	return nil
}