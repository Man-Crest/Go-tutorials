package main

import (
	"fmt"
)

func main() {
	var list [4]string
	// var list = [4]string

	list[0] = "man1"
	list[1] = "man2"
	list[2] = "man3"
	list[3] = "man4"

	fmt.Println(list)
	fmt.Println("length of the list array", len(list))

	var fruits = [3]string{"banana", "mango", "apple"}
	fmt.Println(fruits)

	var arr = [...]int{51, 22, 63, 44}
	fmt.Println(arr)
	fmt.Println("length of the array is", len(arr))
}
