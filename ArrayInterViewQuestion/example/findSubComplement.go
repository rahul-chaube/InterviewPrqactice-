package example

import (
	"fmt"
)

// {1, 7, 2, 5, 9, 0} target 9
func FindSubarrayComplement(input []int, target int) {

	maps := make(map[int]int)
	// maps[0] = -1
	currentSum := 0

	for i, v := range input {

		currentSum += v                   // 1, 8 ,10
		complement := currentSum - target // -8, 0, 1

		fmt.Println(" Step 11 ", i, currentSum, target, complement)

		if index, ok := maps[complement]; ok { // -8 > false , 0 -> false , 1 true 1,2
			fmt.Println(" Index is ", index+1, i, input[index+1:i+1])
			return
		}

		maps[currentSum] = i // 1 > 0 , 8 -> 1 ,

		fmt.Println("Step 2 ", maps)
	}

	fmt.Println("No sub array equal to ", target)
}
