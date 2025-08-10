package level1tasks

import "fmt"

func reverseString(s string) string {
	r := []rune(s)
	res := make([]rune, 0)

	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}

	return string(res)
}

func Task19() {
	str := "главрыба"
	fmt.Printf("%s - %s", str, reverseString(str))
}