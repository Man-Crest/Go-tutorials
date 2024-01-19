package main

import "fmt"

func main() {
	var map1 = make(map[string]string)

	map1["name"] = "man"
	map1["surname"] = "manavadaria"

	fmt.Println(map1)
	fmt.Println(map1["name"])

	var a = map[int]string{1: "man", 2: "nil", 3: "harsh"}
	fmt.Println(a)

	//to delete the map elements
	delete(a, 2)

	//excessing false key
	delete(a, 4) //doesnt care , do not give any error

	//printing false key
	fmt.Println(a[2]) // won'nt panic runs without error , it'll print null values according to the type of value

	fmt.Println(a)

	if _, ok := a[2]; ok {
		fmt.Println(a[2])
	} else {
		fmt.Println("key doesn't exist")
	}
}
