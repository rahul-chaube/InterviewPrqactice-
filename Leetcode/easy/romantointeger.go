package easy

func RomanToInteger(s string) int {
	maps := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	maps['I'] = 1
	maps['V'] = 5
	maps['X'] = 10
	maps['L'] = 50
	maps['C'] = 100
	maps['D'] = 500
	maps['M'] = 1000
	data := 0
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			data += maps[runes[i]]
		} else {
			if maps[runes[i]] < maps[runes[i+1]] {
				data += maps[runes[i+1]] - maps[runes[i]]
				i++
			} else {
				data += maps[runes[i]]
			}
		}
	}
	return data
}
