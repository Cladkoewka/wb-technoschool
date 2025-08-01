package level1tasks

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// блокируем выполнение
	}
}

func Task25() {
	fmt.Println("Before sleep")
	sleep(3 * time.Second)
	fmt.Println("After sleep")
}