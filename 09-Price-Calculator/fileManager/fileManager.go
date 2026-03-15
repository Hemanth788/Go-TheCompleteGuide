package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type SFileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm SFileManager) ReadResultFromInput() ([]string, error) {
		file, err := os.Open(fm.InputFilePath)

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

func (fm SFileManager) WriteResultToOutput(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create a file")
	}

	time.Sleep(3 * time.Second)

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	
	return nil
}

func New(inputPath string, outputPath string) SFileManager {
	return SFileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}