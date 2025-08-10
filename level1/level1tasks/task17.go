package level1tasks

import "fmt"

func binarySearch(arr []int, find int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right - left)/2

		if arr[mid] == find {
			return mid
		} else if arr[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Task17() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 10}
	fmt.Println(binarySearch(arr, 7))
	fmt.Println(binarySearch(arr, 2))
	fmt.Println(binarySearch(arr, 5))
}