package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	PrintTime()
}

// PrintTime output current time to std at format "Time: 15:04:05"
func PrintTime() {
	timeNow, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Time:", timeNow.Format("15:04:05"))
}
