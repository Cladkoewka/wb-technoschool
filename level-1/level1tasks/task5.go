package level1tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func Task5() {
	duration := 5 
	ch := make(chan int)

	timer := time.After(time.Duration(duration) * time.Second)

	// запись значений
	go func() {
		for {
			select {
			case <-timer:
				close(ch)
				return
			default:
				ch <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// чтение значений
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel is closed!")
				return
			}
			fmt.Printf("Get value: %d\n", val)
		case <-timer:
			fmt.Println("Time is over")
			return
		}
	}
}