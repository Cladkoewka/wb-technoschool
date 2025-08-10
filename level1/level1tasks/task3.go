package level1tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func _Work(workersCount int, inChan <-chan int) {
	for i := 0; i < workersCount; i++ {
		go func(workerNum int) {
			for val := range inChan {
				fmt.Printf("Worker %d, value %d\n", workerNum, val)
			}
		}(i)
	}
}

func _TestWorkers() {
	ch := make(chan int)
	workersCount := 5

	// Бесконечная запись рандомных значений в канал
	go func() {
		for {
			ch <- rand.Intn(1000)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Запуск воркеров
	_Work(workersCount, ch)

	// Блокировка main
	select {}
}
