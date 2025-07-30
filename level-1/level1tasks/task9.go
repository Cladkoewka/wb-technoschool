package level1tasks

import (
	"fmt"
	"sync"
)

func generateNumbers(numbers []int, outCh chan<- int) {
	for _, num := range numbers {
		outCh <- num
	}
	close(outCh)
}

func processNumbers(inCh <-chan int, outCh chan<- int) {
	for num := range inCh {
		newNum := num * 2
		outCh <- newNum
	}
	close(outCh)
}

func Task9() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsCh := make(chan int)
	resultCh := make(chan int)

	var wg sync.WaitGroup

	// Горутина для генерации чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		generateNumbers(nums, numsCh)
	}()

	// Горутина для обработки чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		processNumbers(numsCh, resultCh)
	}()

	// Горутина для чтения
	go func() {
		for res := range resultCh {
			fmt.Println(res)
		}
	}()

	wg.Wait()
}
