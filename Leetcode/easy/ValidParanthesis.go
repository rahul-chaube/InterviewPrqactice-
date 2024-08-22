package easy

func IsValid(str string) bool {
	slice := []rune{}
	for _, v := range str {
		switch v {
		case '{', '[', '(':
			slice = append(slice, v)
		case '}', ']', ')':
			if len(slice) == 0 {
				return false
			}
			lastelement := slice[len(slice)-1]
			if lastelement == '[' && v == ']' || lastelement == '{' && v == '}' || lastelement == '(' && v == ')' {
				slice = slice[:len(slice)-1]
			} else {
				return false
			}

		}
	}
	return len(slice) == 0
}
