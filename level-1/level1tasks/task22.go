package level1tasks

import (
	"fmt"
	"math/big"
)

func Task22() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("123456789", 10)
	b.SetString("234567890", 10)

	// сумма
	sum := new(big.Int).Set(a)
	sum.Add(sum, b)
	fmt.Printf("Sum: %s\n", sum.String())

	// разность
	diff := new(big.Int).Set(b)
	diff.Sub(diff, a)
	fmt.Printf("Diff: %s\n", diff.String())

	// умножение
	mul := new(big.Int).Set(a)
	mul.Mul(mul, b)
	fmt.Printf("Mul: %s\n", mul.String())

	// деление
	div := new(big.Int).Set(b)
	div.Div(div, a)
	fmt.Printf("Div: %s\n", div.String())
}