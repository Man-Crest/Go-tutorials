package main

import (
	"fmt"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{} // this are pointers

func main() {

	Wg.Add(1) // number of routins or additional threads are used
	go sheep()
	lion()

	Wg.Wait() // waits for other go routines to complete then terminate main thread
}

func sheep() {
	for i := 1; i <= 10; i++ {
		fmt.Println("sheep")
		time.Sleep(1 * time.Second)
	}
	Wg.Done() //removes thread from wait group and stop wait
}
func lion() {
	for i := 1; i <= 4; i++ {
		fmt.Println("lion")
		time.Sleep(1 * time.Second)
	}
}
