package main

import (
	"fmt"
	"sync"
)

var Wg = sync.WaitGroup{}

func main() {
	// channel := make(chan int)

	// Wg.Add(2)
	// go foo(channel)
	// // fmt.Println("channel data:", <-channel)
	// go temp(channel)
	// // fmt.Println("channel data:", <-channel)

	// Wg.Wait()

	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for chann := range ch {
		fmt.Println("data:", chann)

	}

}

func foo(channel chan int) {
	channel <- 5
	Wg.Done()
}

func temp(channel chan int) {
	channel <- 6
	Wg.Done()
}

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())

// 	var counter int64 = 0

// 	const gs = 100
// 	var wg sync.WaitGroup
// 	// var mu sync.Mutex
// 	wg.Add(gs)

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			// mu.Lock()
// 			// v := counter
// 			atomic.AddInt64(&counter, 1)
// 			time.Sleep(time.Millisecond * 100)
// 			runtime.Gosched()
// 			fmt.Println("counter at atomic:", atomic.LoadInt64(&counter))
// 			// v++
// 			// counter = v
// 			// mu.Unlock()
// 			wg.Done()
// 		}()
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	}
// 	wg.Wait()
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	fmt.Println("count:", counter)
// 	fmt.Println("time", time.Since(start))
// }
