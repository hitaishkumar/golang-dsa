package array

import (
	"math"
)

func FindLargest(input []int) int {
	min := math.MinInt

	for _, val := range input {
		min = max(min, val)
	}

	return min

}
func FindSecondLargest(input []int) int {
	Largest := input[0]
	SecondLargest := input[0]

	for _, val := range input {
		if val > Largest {
			SecondLargest = Largest
			Largest = val
		}

	}

	return SecondLargest

}
func IsArraySorted(input []int) bool {

	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}

	return true

}
func FindUnique(input []int) []int {

	hash := make(map[int]int, 0)
	ans := []int{}
	for i := 1; i < len(input); i++ {
		if hash[input[i]] > 0 {

			hash[input[i]]++

		} else {

			hash[input[i]] = 1
		}
	}

	for key := range hash {
		ans = append(ans, key)
	}

	return ans

}
