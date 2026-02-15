package array

import (
	"fmt"
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
func RotateArray(input []int, amount int) []int {

	steps := amount % len(input)
	store := []int{}
	ans := []int{}

	for idx, val := range input {

		if (idx + 1) <= steps {
			store = append(store, val)
		} else {
			ans = append(ans, val)
		}
	}

	return append(ans, store...)

}

func MoveZerosToEnd(input []int) []int {

	z, nz := 0, 0

	for nz < len(input) {

		if input[nz] != 0 {

			input[z], input[nz] = input[nz], input[z]
			nz++
			z++

		} else {
			nz++
		}

		fmt.Printf("input: %v\n", input)
	}

	return input
}

func UnionOfTwoSorted(inputa, inputb []int) []int {

	left, right := 0, 0
	ans := []int{}

	addToAns := func(val int) {
		if len(ans) == 0 || val != ans[len(ans)-1] {
			ans = append(ans, val)
		}
	}

	for left < len(inputa) && right < len(inputb) {

		if inputa[left] < inputb[right] {
			addToAns(inputa[left])
			left++

		} else if inputa[left] > inputb[right] {
			addToAns(inputa[right])
			right++

		} else {
			addToAns(inputa[right])
			left++
			right++
		}
	}

	for left < len(inputa) {
		addToAns(inputa[left])
		left++
	}
	for right < len(inputb) {
		addToAns(inputb[right])
		right++
	}

	return ans
}
