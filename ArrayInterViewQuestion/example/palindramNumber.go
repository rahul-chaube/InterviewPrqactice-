package example

import (
	"strconv"
)

func IsPalindromNumber(number int) bool {
	convertTostring := strconv.Itoa(number)

	for i := 0; i < len(convertTostring)/2; i++ {
		if convertTostring[i] != convertTostring[len(convertTostring)-1-i] {
			return false
		}
	}
	return true
}
