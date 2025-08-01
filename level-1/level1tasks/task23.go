package level1tasks

import "fmt"

func removeElement(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}

	copy(s[i:], s[i+1:]) // скопировать эл-ты после i на место i
	s = s[:len(s)-1] // уменьшить длину на 1

	return s
}

func Task23() {
	s := []int{1,2,3,4,5,6,7,8,9}

	fmt.Printf("%v\n", s) 
	s = removeElement(s, 2)
	fmt.Printf("%v\n", s) 
}
