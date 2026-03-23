package main

import (
	d "dsa/BS/1D"
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
	// fmt.Println(array.FindUnique([]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}))
	// fmt.Println(array.RotateArray([]int{2, 4, 5, 6}, 5))
	// fmt.Println(array.MoveZerosToEnd([]int{1, 0, 2, 3, 0, 4, 0, 1}))
	// fmt.Println(array.UnionOfTwoSorted([]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 4, 5}))
	// fmt.Println(array.FindMissingNumber([]int{8, 2, 4, 5, 3, 7, 1}))
	// fmt.Println(array.FindMissingNumberXOR([]int{8, 2, 4, 5, 3, 7, 1}))
	// fmt.Println(array.MaxConsecutive([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	// fmt.Println(array.AppearOnce([]int{5, 4, 4, 1, 2, 1, 2}))
	// fmt.Println(array.LongestSubArraySum([]int{10, 5, 2, 7, 1, 9}, 15))
	// fmt.Println(array.LongestSubArraySum2([]int{1, -1, 1, 1, 1, 1}, 3))
	// fmt.Println(array.MajorityElement_I([]int{2, 2, 1, 1, 1, 2, 2}))

	if isFoundAt, err := d.Search([]int{9, 13, 20, 24, 46, 52}, 20); err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Println("isFoundAt -> ", isFoundAt)
	}

}
