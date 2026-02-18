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

func MajorityElement_I(nums []int) int {

	numMap := map[int]int{nums[0]: 1}
	n := len(nums)

	for _, val := range nums[1:] {
		_, ok := numMap[val]
		if ok {
			numMap[val]++
		} else {
			numMap[val] = 1
		}
	}

	for k, v := range numMap {
		fmt.Println(n, n/2)
		if v > (n / 2) {
			return k
		}
	}
	fmt.Printf("numMap: %v\n", numMap)
	return 0
}

// Ideally we should use Boyer-Moore Majority Voting Solution, Next Iteration we will do that
func MajorityElement_II(nums []int) []int {

	freq := make(map[int]int)
	n := len(nums)
	ans := []int{}

	for _, val := range nums {
		freq[val]++
	}

	for k, v := range freq {
		if v > (n / 3) {
			ans = append(ans, k)
		}
	}

	return ans

}
