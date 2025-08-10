package level1tasks

import (
	"fmt"
	"strings"
)

func isSymbolsUnique(s string) bool {
	m := make(map[rune]struct{})
	r := []rune(s)

	for _, symbol := range r {
		lowCaseSymbol := rune(strings.ToLower(string(symbol))[0])
		if _, ok := m[lowCaseSymbol]; ok {
			return false
		} else {
			m[lowCaseSymbol] = struct{}{}
		}
	}

	return true
}

func Task26() {
	fmt.Println(isSymbolsUnique("abcd"))
	fmt.Println(isSymbolsUnique("abCdefAaf"))
	fmt.Println(isSymbolsUnique("aabcd"))
}