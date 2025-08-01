package level1tasks

import "fmt"


// Интерфейс 
type Transport interface {
	Move()
	About()
}

// Структура, реализующая интерфейс
type Car struct {
	Model string
	Color string
}

func (c Car) Move() {
	fmt.Println("Car is riding")
}

func (c Car) About() {
	fmt.Printf("Car %s is %s color\n", c.Model, c.Color)
}

// Структура не реализующая интерфейс
type Person struct {
	Age int
	Name string
}

func (p Person) About() {
	fmt.Printf("%s is %d y.o. \n", p.Name, p.Age)
}

// Адаптер для Person

type PersonAdapter struct {
	person Person
}

func (p PersonAdapter) About() {
	p.person.About()
}

func (p PersonAdapter) Move() {
	fmt.Println("Person isn't a transport, it can't move")
}

func Task21() {
	// Использование Car
	var transport Transport = Car{Model: "BMW", Color: "Black"}
	transport.About()
	transport.Move()

	// Использование Person через адаптер
	person := Person{Name: "Alex", Age: 21}
	personAdapter := PersonAdapter{person: person}
	transport = personAdapter
	transport.About()
	transport.Move()
}