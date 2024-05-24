package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		err = file.Close()
		if err != nil {
			return nil, err
		}
		log.Fatal("Error: ", scanner.Err())

	}
	_ = file.Close()
	return lines, nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		_ = file.Close()
		return errors.New("failed to encode data")
	}
	err = file.Close()
	return nil
}
