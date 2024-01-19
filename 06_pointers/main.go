package main

import (
	"fmt"
)

func main() {
	// var ptr *int
	// fmt.Println(ptr)

	var name string = "man"

	ptr := &name

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Printf("this is the type of ppointer %T \n", ptr)
	fmt.Printf("this is the type of ppointer %T \n", *ptr)

	*ptr += " man"
	fmt.Println(*ptr)
}
