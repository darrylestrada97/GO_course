package conversion

import (
	"errors"
	"log"
	"strconv"
)

func StringToFloats(s []string) ([]float64, error) {
	var floats []float64
	for _, line := range s {
		floatVal, err := strconv.ParseFloat(line, 64)

		if err != nil {
			log.Fatal("Error: ", err)
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
