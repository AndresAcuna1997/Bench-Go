package fileUtils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultValue = 1000

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("No file found")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Faile to parse the stored value")
	}

	return value, nil
}
