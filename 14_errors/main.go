package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	if err != nil {
		fmt.Println("handel the error")
	} else {
		fmt.Println("good to go")
	}

	err1 := DoSomething()
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("good to go")
	}

	ans, err2 := divide(1, 0)
	if err2 == nil {
		fmt.Printf("returned value is %v", ans)
	} else {
		fmt.Printf("you get an error : %v", err2)
	}

}

func foo() error {
	return nil
}
func DoSomething() error {
	return errors.New("something didn't work")
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by 0 value", a)
	}
	return a / b, nil
}
