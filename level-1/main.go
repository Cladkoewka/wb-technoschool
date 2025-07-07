package main

import (
	"fmt"

	"github.com/Cladkoewka/wb-technoschool/level-1/level1tasks"
)

func main() {
	action := level1tasks.Action{
		Human: level1tasks.Human{
			Name:    "Simon",
			Surname: "Riley",
			Age:     21,
		},
		Position: "Programmer",
	}

	action.SayHello()
	fmt.Println(action.BirthYear())
	fmt.Println(action.About())
}
