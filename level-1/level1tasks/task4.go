package level1tasks

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Work(ctx context.Context, workersCount int, inChan <-chan int) {
	wg := sync.WaitGroup{}
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func(workerNum int) {
			defer wg.Done()
			for {
				select {
				case val, ok := <-inChan:
					if !ok {
						fmt.Printf("Worker %d, channel is closed\n", workerNum)
						return // Если канал закрыт завершаем работу
					}
					fmt.Printf("Worker %d, value %d\n", workerNum, val)
				case <-ctx.Done():
					fmt.Printf("Worker %d, ctx done\n", workerNum)
					return // Если сигнал отмены, завершаем работу
				}
			}
		}(i)
	}

	wg.Wait()
}

func generateRandomNumbers(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch) // Закрываем канал при отмене контекста
			return
		default:
			ch <- rand.Intn(1000)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestWorkers() {
	ch := make(chan int)
	workersCount := 5

	// создаем контекст
	ctx, cancel := context.WithCancel(context.Background())

	// Создаем канал с сигналом о прерывании
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// Когда в канале появляется значение, вызываем отмену контекста
	go func() {
		<-signalChan
		cancel()
	}()

	// Бесконечная запись рандомных значений для обработки
	go generateRandomNumbers(ctx, ch)

	// Запуск воркеров
	Work(ctx, workersCount, ch)

	fmt.Println("Work finished")
}
