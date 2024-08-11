package example

import "sort"

func MakeArrayZigZagArray(input []int) []int {

	sort.Ints(input)

	for i := 1; i < len(input); i += 2 {
		input[i], input[i+1] = input[i+1], input[i]

	}
	return input
}
