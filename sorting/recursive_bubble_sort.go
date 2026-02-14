package sorting

func RecursiveBubbleSort(input []int) []int {
	n := len(input)

	for i := 0; i < n-1; i++ {

		for j := i; j < n-1-i; j++ {

			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}

		}

	}

	return input
}
