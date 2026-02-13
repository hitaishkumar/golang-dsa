package sorting

import "fmt"

// merge two sorted array
func merge(a, b []int) []int {

	ans := make([]int, 0, len(a)+len(b))

	left, right := 0, 0

	for left < len(a) && right < len(b) {

		if a[left] <= b[right] {
			ans = append(ans, a[left])
			left++
		} else {
			ans = append(ans, b[right])
			right++
		}
	}

	ans = append(ans, a[left:]...)
	ans = append(ans, b[right:]...)

	fmt.Println(ans)

	return ans
}

func MergeSort(input []int) []int {

	if len(input) == 1 {
		return input
	}

	mid := len(input) / 2

	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])

	return merge(left, right)

}
