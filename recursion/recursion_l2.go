package main

func reverseArr(arr *[]int, left, right int) {

	if left >= right {
		return
	}
	// Dereference the pointer using (*arr) to access indices
	(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
	reverseArr(arr, left+1, right-1)
}
func reverseArrWithOneParam(arr *[]int, idx int) {

	left := idx
	right := len(*arr) - left - 1

	if left >= right {
		return
	}
	// Dereference the pointer using (*arr) to access indices
	(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
	reverseArrWithOneParam(arr, idx+1)
}

func allIndicesInArray(arr *[]int, idx int, val int, fsf int) []int {

	if idx == len(*arr) {
		return make([]int, fsf)
	}
	if (*arr)[idx] == val {
		ans := allIndicesInArray(arr, idx+1, val, fsf+1)
		ans[fsf] = idx
		return ans
	} else {
		ans := allIndicesInArray(arr, idx+1, val, fsf)
		return ans
	}

}

func getSubsequence(input string) []string {
	if len(input) == 1 {
		return []string{"", input}
	}

	ss := getSubsequence(input[1:])
	myAns := make([]string, 0)

	firstChar := input[0:1]

	for _, value := range ss {

		myAns = append(myAns, value)
		myAns = append(myAns, firstChar+value)

	}

	return myAns
}

func getStairsPath(start, end int) []int {

	if start <= end {

		return []int{1}
	}

	// path1 := getStairsPath(start-1, end)
	// path2 := getStairsPath(start-2, end)
	// path3 := getStairsPath(start-3, end)

	return []int{1}
}
