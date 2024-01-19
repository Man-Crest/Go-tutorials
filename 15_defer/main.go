package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	defer temp()
}

func temp() {
	defer fmt.Println("hello1")
	defer fmt.Println("hello2")
	defer fmt.Println("hello3")
	defer fmt.Println("hello4")
}
