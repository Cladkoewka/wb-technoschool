package level1tasks

import "fmt"

func Task13() {
	a := 8
	b := 4

	fmt.Printf("(Before switch) a: %d, b: %d\n", a, b)

	a = a + b // сумма a + b
	b = a - b // (a + b) - b = a
	a = a - b // (a + b) - a = b

	fmt.Printf("(After switch) a: %d, b: %d\n", a, b)
}