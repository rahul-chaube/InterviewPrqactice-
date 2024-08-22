package easy

func IsPolidromNumber(x int) bool {
	temp := 0
	reverse := x
	for x > 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	return temp == reverse
}
