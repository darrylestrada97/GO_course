package conversion

import (
	"errors"
	"log"
	"strconv"
)

func StringToFloat(s []string) ([]float64, error) {
	for lineIndex, line := range s {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			log.Fatal("Error: ", err)
			return nil, errors.New("failed to convert string to float")
		}
		prices[lineIndex] = floatPrice
	}
}
