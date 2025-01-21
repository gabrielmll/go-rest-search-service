package service

// BinarySearchValue performs a binary search on a sorted slice of integers to find a target value.
// If the exact target is not found, it checks for the closest value within a 10% margin.
//
// Parameters:
// - numbers: A sorted slice of integers to search within.
// - target: The target value to search for.
//
// Returns:
// - int: The index of the target or closest value within the 10% margin.
// - int: The value at the returned index, or -1 if no match is found.
// - string: An empty string if a match is found, or an error message if no value is within the margin.
func BinarySearchValue(numbers []int, target int) (int, int, string) {
	low, high := 0, len(numbers)-1
	for low <= high {
		mid := (low + high) / 2
		if numbers[mid] == target {
			return mid, numbers[mid], ""
		} else if numbers[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// 10% margin logic
	closest := findClosest(numbers, target, low, high)
	if closest != -1 {
		return closest, numbers[closest], ""
	}

	return -1, -1, "Value not found within acceptable margin"
}

// findClosest searches for the closest number to the target within a 10% margin in a sorted slice of integers.
// It checks the low and high indices to determine if either of them falls within the acceptable range of values.
//
// Parameters:
// - numbers: A sorted slice of integers to search within.
// - target: The target value to find the closest match for.
// - low: The lower index to check.
// - high: The higher index to check.
//
// Returns:
// - The index of the closest number within the 10% margin if found.
// - -1 if no number within the margin is found.
func findClosest(numbers []int, target, low, high int) int {
	margin := target / 10
	lowerBound := target - margin
	upperBound := target + margin

	if low >= 0 && low < len(numbers) && numbers[low] >= lowerBound && numbers[low] <= upperBound {
		return low
	}
	if high >= 0 && high < len(numbers) && numbers[high] >= lowerBound && numbers[high] <= upperBound {
		return high
	}

	return -1
}
