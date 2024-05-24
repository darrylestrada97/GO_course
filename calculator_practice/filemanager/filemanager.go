package filemanager

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error: ", err)
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
			fmt.Print("Error: ", err)
		}
		log.Fatal("Error: ", scanner.Err())

	}
	_ = file.Close()
	return lines
}
