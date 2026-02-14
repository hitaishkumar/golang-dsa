package main

import (
	"dsa/array"
	"fmt"
)

func main() {

	// printNameNTimes("Hitaish", 5)

	// printNum(1, 10)

	// printNumReverse(1, 10)

	// fmt.Println(sumNumbers(1, 10))
	// fmt.Println(factorial(10))

	// nums := []int{1, 2, 3, 4, 5}
	// reverseArr(&nums, 0, len(nums)-1)
	// fmt.Println(nums)
	// nums := []int{1, 2, 3, 4, 5}
	// reverseArrWithOneParam(&nums, 1)
	// fmt.Println(nums)

	// input := "abcdef"
	// results := getSubsequence(input)

	// fmt.Printf("Input: %s\n", input)
	// fmt.Printf("Total Subsequences: %d\n", len(results))
	// fmt.Printf("Results: %v\n", results)

	// // Optional: Print them one by one for better clarity
	// fmt.Println("\nDetailed List:")
	// for i, s := range results {
	// 	fmt.Printf("%d: '%s'\n", i+1, s)
	// }

	// Example 1:
	// Input: N = 6, array[] = {13,46,24,52,20,9}

	// fmt.Println(sorting.SelectionSort([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(sorting.BubbleSort([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(sorting.RecursiveBubbleSort([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(sorting.MergeSort([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(array.FindLargest([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(array.FindSecondLargest([]int{13, 46, 24, 52, 20, 9}))
	// fmt.Println(array.IsArraySorted([]int{-2, 1}))
	fmt.Println(array.FindUnique([]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}))
}
