package easy

func SumOfTwoNumber(arr []int, target int) []int {
	data := make(map[int]int)
	for i, v := range arr {
		complement := target - v

		if data, isExists := data[complement]; isExists {
			return []int{data, i}
		}
		data[v] = i
	}
	return []int{0, 0}
}
