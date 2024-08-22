package easy

import "fmt"

func DiffString(str1, str2 string) string {
	catch := make(map[rune]int)

	for _, v := range str1 {
		catch[v]++
	}

	for _, v := range str2 {
		catch[v]--
	}
	s := []rune{}
	for i, v := range catch {
		if v != 0 {
			s = append(s, i)
		}
	}
	fmt.Println(" ", catch)
	return string(s)
}
