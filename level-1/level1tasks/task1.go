package level1tasks

import (
	"fmt"
	"time"
)

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h Human) SayHello() {
	fmt.Printf("Hello, my name is %s %s, i'm %d years old!\n", h.Name, h.Surname, h.Age)
}

func (h Human) BirthYear() int {
	return time.Now().Year() - h.Age
}

type Action struct {
	Human
	Position string
}

func (a Action) About() string {
	return fmt.Sprintf("I'm %s, my position is %s", a.Name, a.Position)
}