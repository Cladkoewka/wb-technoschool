package level1tasks

import "fmt"

func reverseWords(words string) string {
	r := []rune(words)

	// разворот всей строки
	reverse(r, 0, len(r)-1)

	// разворот слов
	start := 0
	for i := 0; i < len(r); i++ {
		if i == len(r) || r[i] == ' '{
			reverse(r, start, i-1)
			start = i + 1
		}
	}

	return string(r)
}

func reverse(r []rune, start, end int) {
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}

func Task20() {
	str := "snow dog sun"
	fmt.Printf("%s - %s\n", str, reverseWords(str))
	str = "s"
	fmt.Printf("%s - %s\n", str, reverseWords(str))
	str = ""
	fmt.Printf("%s - %s\n", str, reverseWords(str))
}