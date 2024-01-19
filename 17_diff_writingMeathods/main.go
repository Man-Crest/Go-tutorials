package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//usinf bufio.NewWriter
	file, err := os.Create("./newFile1.txt")
	checkNil(err)

	content := "hgrthathqwerdsz"
	writer := bufio.NewWriter(file)
	writer.Write([]byte(content))
	writer.Flush()
	file.Close()

	file1, err := os.Open("./newFile2.txt")
	checkNil(err)
	reader := bufio.NewReader(file1)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
