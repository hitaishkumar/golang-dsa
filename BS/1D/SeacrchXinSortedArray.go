package d

import (
	"cmp"
	"errors"
)

func Search[T cmp.Ordered](input []T, target T) (int, error) {
	low, high := 0, len(input)-1
	for low < high {
		mid := (low + high) / 2
		if input[mid] == target {
			return mid, nil
		}
		if target < input[mid] {
			high = mid - 1
		}
		if target > input[mid] {
			low = mid + 1
		}
	}
	return -1, errors.New("Not Found")
}

// Example 2:
//
//	Input Format: N = 5, arr[] = {3,5,8,15,19}, x = 9
//	Result: 3
//	Explanation: Index 3 is the smallest index such that arr[3] >= x.
func SearchLowerBound[T cmp.Ordered](input []T, target T) int {
	low, high := 0, len(input)-1
	for low < high {
		mid := (low + high) / 2
		if target < input[mid] {
			high = mid
		}
		if target > input[mid] {
			low = mid + 1
		}
	}

	return low
}

// Example 2:
//
//	Input Format: N = 6, arr[] = {3,5,8,9,15,19}, x = 9
//										  ^
//	Result: 4
//	Explanation: Index 4 is the smallest index such that arr[4] > x.
func SearchUpperBound[T cmp.Ordered](input []T, target T) int {
	low, high := 0, len(input)-1
	for low < high {
		mid := (low + high) / 2
		if target < input[mid] {
			high = mid + 1
		}
		if target > input[mid] {
			low = mid + 1
		}
	}

	return low
}
