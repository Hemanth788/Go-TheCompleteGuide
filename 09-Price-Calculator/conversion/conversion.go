package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := []float64{}
	for _, string := range strings {
		floatVal, err := strconv.ParseFloat(string, 64)
		if err != nil {
			fmt.Println("Failed to convert the string to a float64")
			fmt.Println(err)
			return nil, errors.New("Failed to convert the string to a float64")
		} 
		floats = append(floats, floatVal)
	}
	return floats, nil
}