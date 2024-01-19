package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=css&username=man"

func main() {
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println("scheme", result.Scheme)
	fmt.Println("path", result.Path)

	fmt.Println("host", result.Host)
	fmt.Println("port:", result.Port())
	fmt.Println("raw", result.RawQuery)
	fmt.Println("query", result.Query())

	for _, val := range result.Query() {
		fmt.Println(val)
	}

	//must have to pass reference of url.URL using &
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "learn",
		// RawQuery: "coursename=css&username=man",
	}

	mainUrl := partsOfUrl.String()

	fmt.Println(mainUrl)
}
