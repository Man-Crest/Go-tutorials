package calculator

import "fmt"

func Add(a int, b int) int {

	fmt.Println(Val) // we can use val wether it is exported or not because val is initialised in same package that is claculator
	return a + b
}

// func sub(a int, b int) int {
// 	return a - b
// }
// func mul(a int, b int) int {
// 	return a * b
// }
