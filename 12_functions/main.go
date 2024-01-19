package main

import "fmt"

func main() {
	// temp()
	// result := greeter(2, 3)
	// fmt.Println("result :", result)
	res1, res2 := proadder(1, 1, 1, 1, 1, 1)
	// res1, _:= proadder(1, 1, 1, 1, 1, 1)

	fmt.Println(res1, res2)

	//ANNONIMUS functions

	func() {
		fmt.Println("annaonimus")
	}()

	fmt.Println(first(5, returnParams))

}

func returnParams(i int) int {
	return i
}

//function as a parameter

func first(a int, f func(int) int) int {
	return f(a)
}

func greeter(val1 int, val2 int) int {
	return val1 + val2
}

func temp() {
	fmt.Println("this is a temporery fuction")
}

func proadder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "hello"
}
