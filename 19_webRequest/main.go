package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var url string = "https://lco.dev"

func main() {
	responce, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Printf("this is the type of responce : %T \n", responce)
		data, err := io.ReadAll(responce.Body)

		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(string(data))
	}
	defer responce.Body.Close()
}
