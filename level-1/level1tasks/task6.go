package level1tasks

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func exitByCondition() {
	isStop := false

	go func() {
		for !isStop {
			fmt.Println("Goroutine is working")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Goroutine is stopped")
	}()

	time.Sleep(2 * time.Second)
	isStop = true
	time.Sleep(500 * time.Millisecond)
}

func exitByChannel() {
	ch := make(chan struct{})

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Goroutine is stopped")
				return
			default:
				fmt.Println("Goroutine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ch <- struct{}{}
	time.Sleep(500 * time.Millisecond)
}

func exitByContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine is stopped")
				return
			default:
				fmt.Println("Goroutine is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func exitByGoexit() {
	go func() {
		i := 0
		for {
			fmt.Println("Goroutine is working")
			time.Sleep(500 * time.Millisecond)
			i++
			if i > 5 { // пример условия выхода
				runtime.Goexit()
			}
		}
	}()

	time.Sleep(3 * time.Second)
}

func Task6() {
	fmt.Println("===exitByCondition===")
	exitByCondition()

	fmt.Println("===exitByChannel===")
	exitByChannel()

	fmt.Println("===exitByContext===")
	exitByContext()

	fmt.Println("===exitByGoexit===")
	exitByGoexit()
}
