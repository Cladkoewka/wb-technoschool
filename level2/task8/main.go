package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	PrintTime()
}

func PrintTime() {
	timeNow, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения: ", err)
		os.Exit(1)
	}

	fmt.Println("Время:", timeNow.Format("15:04:05"))
}
