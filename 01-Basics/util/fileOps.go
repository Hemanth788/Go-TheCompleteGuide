package util // in the folder inside the project with the same name

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	
	if err != nil {
		return 0, errors.New("Failed to find/read file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("There is non-float content in the file")
	}


	return balance, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value);
	os.WriteFile(fileName, []byte(valueText), 0644)
}