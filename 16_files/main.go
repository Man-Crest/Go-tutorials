package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")

	content := "hello this is a string that will be saved into new file"

	file, err := os.Create("./myNewFile.txt")

	checkNil(err)

	length, err := io.WriteString(file, content)

	checkNil(err)

	fmt.Println("length is ", length)

	// readFile("./myNewFile.txt")

	defer file.Close()

	//file is exist or not !!!
	isFileExist("./myNewFile.txt")
}

func isFileExist(filename string) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Println("file doesn't exist")
	} else {
		fmt.Println("file exists")
	}
}

func readFile(filename string) {
	dataBytes, err := os.ReadFile(filename)
	checkNil(err)
	fmt.Println("your file is this in bytes : ", dataBytes)
	fmt.Println("your file is this in string : ", string(dataBytes))
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
