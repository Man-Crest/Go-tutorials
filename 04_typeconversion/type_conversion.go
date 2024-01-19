package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 5
	fmt.Printf("type of num before convirsion %T \n", num)

	num2 := float32(num)
	fmt.Printf("type of num2 before convirsion %T \n", num2)

	//string to int

	a := "43"
	fmt.Printf("type of a before convirsion %T \n", a)
	data, _ := strconv.Atoi(a)
	fmt.Printf("type of data after convirsion %T \n", data)

	// int to string
	b := 45
	fmt.Printf("type of a before convirsion %T \n", b)
	data2 := strconv.Itoa(b)
	fmt.Printf("type of data after convirsion %T \n", data2)

}
