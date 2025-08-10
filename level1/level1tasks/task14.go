package level1tasks

import "fmt"

func concreteType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Printf("%v is integer\n", t)
	case string:
		fmt.Printf("%v is string\n", t)
	case bool:
		fmt.Printf("%v is bool\n", t)
	case chan int:
		fmt.Printf("%v is chan int\n", t)
	default:
		fmt.Printf("%v is other type\n", t)
	}
}

func Task14() {
	concreteType(13)
	concreteType("Hello")
	concreteType(true)
	concreteType(make(chan int))
	concreteType(make(chan string))
}
