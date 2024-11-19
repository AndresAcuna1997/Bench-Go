package parse_utils

import (
	"errors"
	"fmt"
	"strconv"
)

func TransformStrinToFloat(s string) (float64, error) {

	float, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, errors.New("error parsin the string into a float")
	}

	return float, nil
}

func TransformFloatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
