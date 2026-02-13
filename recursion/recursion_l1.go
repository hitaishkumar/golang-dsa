package main

import "fmt"

func printNameNTimes(name string, N int) {
	if N <= 0 {
		return
	}

	fmt.Println(name)
	printNameNTimes(name, N-1)
}

func printNum(Start, End int) {

	if Start > End {
		return
	}
	fmt.Println(Start)
	printNum(Start+1, End)
}

func printNumReverse(Start, End int) {

	if End < Start {
		return
	}
	fmt.Println(End)
	printNumReverse(Start, End-1)
}

func sumNumbers(start, end int) int {
	if start == end {
		return start
	}
	return start + sumNumbers(start+1, end)
}
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
