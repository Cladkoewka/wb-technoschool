package level2tasks

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func Task3() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
