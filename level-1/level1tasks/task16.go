package level1tasks

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // массив из 0 или 1 эл-та уже отсортирован
	}

	// опорный эл-т
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// срезы для эл-ов меньших, равных и больших опорного
	var less, equal, greater []int
	for _, val := range arr {
		if val < pivot {
			less = append(less, val)
		} else if val > pivot {
			greater = append(greater, val)
		} else {
			equal = append(equal, val)
		}
	}
	
	// рекурсивный вызов сортировки подмассивов
	sortedLess := quickSort(less)
	sortedGreater := quickSort(greater)

	// Объединение 
	combined := append(sortedLess, equal...)
	result := append(combined, sortedGreater...)

	return result
}

func Task16() {
	arr := []int{3, 6, 8, 10, 1, 2, 1, 4, 3, 7, 2}
	fmt.Printf("Before sort: %v\n", arr)
	fmt.Printf("After sort: %v\n", quickSort(arr))
}