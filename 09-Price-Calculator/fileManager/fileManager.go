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

	defer file.Close() // -> instead of closing files manually when an error occurs or at the end of the function, you defer it, once the file is usable that is, and go takes care of it for you

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Failed to read the file content")
	}

	return lines, nil
}

func (fm SFileManager) WriteResultToOutput(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create a file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
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