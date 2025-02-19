package test

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getFloatToFile(fileName string, defaultValueReturn float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValueReturn, errors.New("File not found")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
