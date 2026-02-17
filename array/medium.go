package array

import "fmt"

func LongestSubArraySum2(input []int, desired_sum int) int {
	curr_sum := 0
	res := 0
	prefix_sum_map := map[int]int{0: 1}

	for _, val := range input {

		curr_sum += val
		diff := curr_sum - desired_sum

		if value, ok := prefix_sum_map[diff]; ok {
			res += value
		}

		prefix_sum_map[curr_sum]++

		fmt.Printf("prefix_sum_map: %v\n", prefix_sum_map)

	}

	return res
}
