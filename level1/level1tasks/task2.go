package level1tasks

import (
	"fmt"
	"sync"
)

func calculateSquares(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go func(val int) {
			defer wg.Done()

			square := v * v
			fmt.Printf("%d^2 = %d\n", v, square)
		}(v)
	}
	wg.Wait()
}

func TestCalculateSquares() {
	arr := []int{2, 4, 6, 8, 10}
	calculateSquares(arr)
}
