package level1tasks

import "fmt"

func Task10() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)

	for _, temp := range temperatures {
		key := int(temp / 10) * 10
		m[key] = append(m[key], temp)
	}

	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
}