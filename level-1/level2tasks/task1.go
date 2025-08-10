package level2tasks

import "fmt"

func Task1() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // срез на основе массива a, последний элемент не включается
	fmt.Println(b) // [77, 78, 79]
}