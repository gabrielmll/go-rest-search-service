package utils

import (
	"fmt"
	"os"
)

// LoadNumbersFile loads the numbers from the file into a slice
func LoadNumbersFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var numbers []int
	for {
		var num int
		_, err := fmt.Fscanf(file, "%d\n", &num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
