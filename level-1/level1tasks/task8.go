package level1tasks

import (
	"errors"
	"fmt"
)

func setIthBit(num *int64, i int, isZero bool) error {
	if i < 0 || i > 63 {
		return errors.New("index i must be in range 0 and 63")
	}

	if isZero {
		*num &^= (1 << i) // Устанавливаем в 0 i-ый бит
	} else {
		*num |= (1 << i) // Устанавливаем в единицу
	}

	return nil
}

func Task8() {
	var num int64 = 5

	fmt.Printf("Num before set: %d\n", num)
	setIthBit(&num, 0, true)
	fmt.Printf("Num after set: %d\n", num)
}