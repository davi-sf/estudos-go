package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(valueData []string) ([]float64, error) {
	valueFloats := make([]float64, len(valueData))

	for i := 0; i < len(valueData); i++ {
		floatPrice, err := strconv.ParseFloat(valueData[i], 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		valueFloats[i] = floatPrice
	}

	return valueFloats, nil

}
