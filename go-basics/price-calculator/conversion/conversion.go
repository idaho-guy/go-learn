package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Faild to convert string to float")
		}
		floats = append(floats, floatPrice)

	}
	return floats, nil
}
