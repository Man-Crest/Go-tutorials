package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("main.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("some error ouccured", err.Error())
	}

	defer file.Close()

	_, err1 := file.WriteString("hello this is man")
	if err1 != nil {
		log.Fatal("some error ouccured", err.Error())
	}
}
