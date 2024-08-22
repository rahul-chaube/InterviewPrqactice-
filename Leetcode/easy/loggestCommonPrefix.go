package easy

import (
	"sort"
)

func findCommonPrefix(str []string) string {

	sort.Strings(str)
	first := str[0]
	last := str[len(str)-1]

	minLength := 0
	if len(first) < len(last) {
		minLength = len(first)
	} else {
		minLength = len(last)
	}
	counter := 0
	for i := 0; i < minLength; i++ {
		if first[i] != last[i] {
			break
		}
		counter++
	}
	if counter == 0 {
		return "-1"
	}
	return first[:counter]

}
