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

func FindMissingNumber(input []int) int {
	n := len(input)
	ap_sum := (n + 1) * (n + 2) / 2 // n(n+1)/2 // n-> number of terms

	full_sum := 0

	for _, val := range input {
		full_sum += val
	}

	return ap_sum - full_sum
}
func FindMissingNumberXOR(input []int) int {
	xor_full := 0
	n := len(input)
	xor_partial := 0

	for _, val := range input {
		xor_partial ^= val
	}
	for i := 1; i <= n+1; i++ {
		xor_full ^= i
	}

	return xor_full ^ xor_partial
}

func MaxConsecutive(input []int) int {

	//{1, 1, 0, 1, 1, 1}

	max_sum := 0
	sum := 0

	for i := 0; i < len(input); i++ {

		if input[i] == 0 {
			sum = 0
		} else {
			sum += 1
		}
		if sum > max_sum {
			max_sum = sum
		}
	}

	return max_sum

}

func AppearOnce(input []int) int {
	xor_val := input[0]
	for i := 1; i < len(input); i++ {
		xor_val ^= input[i]
	}

	return xor_val
}
func LongestSubArraySum(input []int, sum int) int {
	sub_arr := []int{}
	max_len_sub_arr := 0
	sum_arr := 0
	for i := 0; i < len(input); i++ {
		if sum_arr < sum {
			sub_arr = append(sub_arr, input[i])
			sum_arr += input[i]
		}
		fmt.Println(sub_arr)
		fmt.Println(sum_arr)
		if sum_arr == sum {
			if len(sub_arr) > max_len_sub_arr {
				max_len_sub_arr = len(sub_arr)
			}
			sum_arr = 0
			sub_arr = []int{}
		}

	}

	return max_len_sub_arr
}
