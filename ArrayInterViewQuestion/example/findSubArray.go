package example

func FindSubArrayMaxSum(input []int) int {
	outermax := 0
	for i := 0; i < len(input); i++ {
		var innermax = 0
		for j := i; j < len(input); j++ {
			innermax += input[j]

			if innermax > outermax {
				outermax = innermax
			}

		}

	}
	return outermax
}

func KadaneAlorith(input []int) int {
	maxSoFar := 0
	maxendingHere := 0

	for _, v := range input {
		maxendingHere = maxendingHere + v

		if maxSoFar < maxendingHere {
			maxSoFar = maxendingHere
		}

		if maxendingHere < 0 {
			maxendingHere = 0
		}
	}
	return maxSoFar
}

// []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
func KadaneAlorithSubArray(input []int) (int, []int) {
	maxSoFar := input[0]
	maxendingHere := input[0]

	start, end, s := 0, 0, 0

	for i := 1; i < len(input); i++ {
		if input[i] > maxendingHere+input[i] {

			maxendingHere = input[i]

			s = i // Setting starting point
		} else {
			maxendingHere += input[i]

		}

		if maxendingHere > maxSoFar {
			maxSoFar = maxendingHere
			start = s
			end = i
		}

	}
	return maxSoFar, input[start : end+1]
}
