package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	// Read the balance from a file
	valueText, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	value, err := strconv.ParseFloat(string(valueText), 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse float")
	}
	return value, nil
}
func WriteFloatToFile(value float64, fileName string) {
	// Write the balance to a file
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		fmt.Println("Failed to save balance to file")
	}
}
