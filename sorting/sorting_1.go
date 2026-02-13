package sorting

import "fmt"

// Example 1:
// Input: N = 6, array[] = {13,46,24,52,20,9}
// Output: 9,13,20,24,46,52
// Explanation: After sorting the array is: 9, 13, 20, 24, 46, 52

// Example 2:
// Input: N=5, array[] = {5,4,3,2,1}
// Output: 1,2,3,4,5
// Explanation: After sorting the array is: 1, 2, 3, 4, 5

func SelectionSort(input []int) []int {
	for idx, value := range input {

		if idx == len(input)-1 {
			break
		}

		minIdx := findMinIndex(input[idx+1:])
		input[idx], input[minIdx+idx+1] = input[minIdx+idx+1], value
	}
	return input

}

func findMinIndex(input []int) int {
	min := 0

	for idx, value := range input {
		if value < input[min] {
			min = idx
		}
	}

	return min
}

func BubbleSort(input []int) []int {
	// The Objective here to do
	// 1. In each interation , take teh max value to the right most by swapping values till the end index
	// 2. Repeat the same for all value in the array

	n := len(input)

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}

		}

	}

	return input
}

func InsertionSort(input []int) []int {
	n := len(input)
	for i := 0; i < n-1; i++ {

		for j := i; j > 0; j-- {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
		}
		fmt.Println(input)

	}

	return input
}
