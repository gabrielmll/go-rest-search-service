package service

import (
	"testing"
)

// TestBinarySearchValue validates the behavior of the binary search function.
func TestBinarySearchValue(t *testing.T) {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	tests := []struct {
		name          string
		target        int
		expectedIndex int
		expectedValue int
		expectedError string
	}{
		{
			name:          "Exact match found",
			target:        20,
			expectedIndex: 2,
			expectedValue: 20,
			expectedError: "",
		},
		{
			name:          "Within 10% margin (lower bound)",
			target:        22, // 22 is within 10% of 25
			expectedIndex: 2,
			expectedValue: 20,
			expectedError: "",
		},
		{
			name:          "Within 10% margin (upper bound)",
			target:        38, // 40 is within 10% of 38
			expectedIndex: 4,
			expectedValue: 40,
			expectedError: "",
		},
		{
			name:          "Not found within 10% margin",
			target:        5, // No value is within 10% of 5
			expectedIndex: -1,
			expectedValue: -1,
			expectedError: "Value not found within acceptable margin",
		},
		{
			name:          "Out of range (too high)",
			target:        200, // No value is within 10% of 200
			expectedIndex: -1,
			expectedValue: -1,
			expectedError: "Value not found within acceptable margin",
		},
		{
			name:          "Out of range (too low)",
			target:        -10, // No value is within 10% of -10
			expectedIndex: -1,
			expectedValue: -1,
			expectedError: "Value not found within acceptable margin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, value, err := BinarySearchValue(numbers, tt.target)

			if index != tt.expectedIndex {
				t.Errorf("Expected index: %d, got: %d", tt.expectedIndex, index)
			}

			if value != tt.expectedValue {
				t.Errorf("Expected value: %d, got: %d", tt.expectedValue, value)
			}

			if err != tt.expectedError {
				t.Errorf("Expected error: '%s', got: '%s'", tt.expectedError, err)
			}
		})
	}
}

// TestFindClosest ensures this helper function accurately identifies the closest value within a 10% margin.
func TestFindClosest(t *testing.T) {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	tests := []struct {
		name          string
		target        int
		low           int
		high          int
		expectedIndex int
	}{
		{
			name:          "Closest value in range (lower bound)",
			target:        22,
			low:           2,
			high:          3,
			expectedIndex: 2,
		},
		{
			name:          "Closest value in range (upper bound)",
			target:        37,
			low:           3,
			high:          4,
			expectedIndex: 4,
		},
		{
			name:          "No value within margin",
			target:        5,
			low:           0,
			high:          1,
			expectedIndex: -1,
		},
		{
			name:          "Out of range (too high)",
			target:        150,
			low:           10,
			high:          11,
			expectedIndex: -1,
		},
		{
			name:          "Out of range (too low)",
			target:        -10,
			low:           -1,
			high:          0,
			expectedIndex: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := findClosest(numbers, tt.target, tt.low, tt.high)

			if index != tt.expectedIndex {
				t.Errorf("Expected index: %d, got: %d", tt.expectedIndex, index)
			}
		})
	}
}
