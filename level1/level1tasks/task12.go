package level1tasks

import "fmt"

func uniqueWords(words []string) []string {
	m := make(map[string]struct{})

	for _, word := range words {
		m[word] = struct{}{}
	}

	res := make([]string, 0)
	for word := range m {
		res = append(res, word)
	}

	return res
}

func Task12() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("%v\n", uniqueWords(words))
}