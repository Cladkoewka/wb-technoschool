package level1tasks

import "fmt"

func intersection(set1, set2 []int) []int {
	m := make(map[int]struct{})

	// отмечаем в мапе встреченные значения
	for _, val := range set1 {
		m[val] = struct{}{}
	}

	res := make([]int, 0)
	seen := make(map[int]bool) // Если во втором множестве могут дублироваться значения, это поможет избежать дубликатов
	for _, val := range set2 {
		if _, ok := m[val]; ok && !seen[val] {
			res = append(res, val)
			seen[val] = true
		}
	}

	return res
}

func Task11() {
	set1, set2 := []int{1, 2, 3}, []int{2, 3, 4, 2, 2, 3}
	fmt.Printf("%v\n", intersection(set1, set2))
}
