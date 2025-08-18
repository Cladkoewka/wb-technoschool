package level2tasks

import (
	"fmt"
	"sort"
	"strings"
)

// sortRunes - returns string with sorted symbols
func sortRunes(s string) string {
	r := []rune(s)
	sort.Slice(r, func(a,b int) bool {
		return r[a] < r[b]
	})

	return string(r)
}

// findAnagrams - search set of anagrams
func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	firstWords := make(map[string]string)

	for _, word := range words {
		lower := strings.ToLower(word)
		sorted := sortRunes(lower)

		if _, ok := firstWords[sorted]; !ok {
			firstWords[sorted] = lower
		}
		anagrams[sorted] = append(anagrams[sorted], lower)
	}


	result := make(map[string][]string)
	for sorted, group := range anagrams {
		if len(group) > 1 {
			sort.Strings(group)
			result[firstWords[sorted]] = group
		}
	}

	return result
}

func Task11() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	//words := []string{}
	//words := []string{"лепс", "СПЕЛ"}
	anagrams := findAnagrams(words)

	for key, group := range anagrams {
		fmt.Printf("%q: %v\n", key, group)
	}
}

