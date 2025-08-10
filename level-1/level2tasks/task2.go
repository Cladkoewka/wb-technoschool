package level2tasks

import "fmt"

func test() (x int) { // объявление именованной возвращаемой переменной. Изначально 0
  defer func() { // выполняется после завершения функции
    x++
  }()
  x = 1
  return // перед фактическим возвратом Go выполняет все defer-функции
}

func anotherTest() int {
  var x int // объявление переменной
  defer func() { 
    x++
  }()
  x = 1 // присваиваем значение
  return x
}

func Task2() {
  fmt.Println(test()) // 2
  fmt.Println(anotherTest()) // 1
}