package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// decodeJson()
	encodeJson()
}

func encodeJson() {
	p2 := []person{
		{"a", 1, "gj1"},
		{"b", 2, "gj2"},
		{"c", 3, "gj3"},
	}

	// fields of struct must be Capitalized or exported otherwise it will nt work
	val, _ := json.Marshal(p2)
	val2, _ := json.MarshalIndent(p2, "", "\t")
	fmt.Println("output", string(val))
	fmt.Println("output", string(val2))
}

func decodeJson() {
	jsonData := []byte(`
	{
		"name":"John",
		"age":30, 
		"city":"New York"
	}
	`)

	var P1 person
	check := json.Valid(jsonData)
	if check {
		json.Unmarshal(jsonData, &P1)
	} else {
		fmt.Println("data is not in valid syntex")
	}

	fmt.Println("data:", P1)
}
