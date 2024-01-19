package main

import "fmt"

const PublicVariable = "globle"

// by using the capitle latter at initial we can set variable globle.

func main() {
	var username string = "Man Manavadaria"
	fmt.Println(username)
	fmt.Printf("the type of username is : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("type of isloggedin is %T \n", isLoggedin)

	var smallval uint8 = 255
	// var smallval uint16 = 255
	// var smallval uint32 = 255
	// var smallval uint64= 255
	fmt.Println(smallval)
	fmt.Printf("type of smallval is %T \n", smallval)

	var floatVal float32 = 255.3245415315
	fmt.Println(floatVal)
	fmt.Printf("type of floatVal is %T \n", floatVal)

	var floatVal2 float64 = 255.3245415315
	fmt.Println(floatVal2)
	fmt.Printf("type of floatVal2 is %T \n", floatVal2)

	// without assigning values to declaration
	var anotherVar int // default value 0 is added to variable
	fmt.Println(anotherVar)

	//implecit types :- without giving type to variables
	var company = "crest"
	fmt.Println(company)
	// company= 12121 will give error because go did auto assign type according to the value

	// WITHOU USING VAR
	name := "man"
	fmt.Println(name)
	// we can declare variable outside of main func but without var is not allowd, have to use var key word
	// const can not be declared by :=

	fmt.Println(PublicVariable)

	//declaration block
	var (
		product  = "Mobile"
		quantity = 50
		price    = 50.50
		inStock  = true
	)
	fmt.Println(product, quantity, price, inStock)

	name1, name2 := "man", "harsh"
	fmt.Println(name1, name2)

}
