package file

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/tax-calculator/parse_utils"
)

type FileManager struct {
	InputFilePath  string
	OutPutFilePath string
}

func (fm FileManager) ReadDataFile() ([]float64, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Error loading file")
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error loading file")
		fmt.Println(err)
		file.Close()
		return nil, err
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := parse_utils.TransformStrinToFloat(line)

		if err != nil {
			file.Close()
			fmt.Println("Error parsin the string into a float")
			return nil, err
		}

		prices[lineIndex] = floatPrice
	}

	file.Close()
	return prices, nil
}

func (fm FileManager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutPutFilePath)

	if err != nil {
		return errors.New("failed too create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("fail to convert data to JSON")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath,
		outputPath,
	}
}
