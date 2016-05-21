package main

import (
	"bufio"
	"os"
)

func ReadFixtures(path string, size int) []string {
	file, err := os.Open(path)
	HandleError(err)
	defer file.Close()

	result := []string{}
	imported_size := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
		imported_size++
		if size != -1 && imported_size >= size {
			break
		}
	}

	return result
}

func ReadCabs(path string, size int) []*Cab {
	result := []*Cab{}
	coordinates := ReadFixtures(path, size)
	for i := 0; i < len(coordinates); i++ {
		result = append(result, NewCabFromCoordinates(coordinates[i], true))
	}
	return result
}
