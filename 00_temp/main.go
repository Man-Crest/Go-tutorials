// write methods
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	data := []byte("Hello, Golang!")
// 	n, err := file.Write(data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Wrote %d bytes to file.\n", n)
// }

// bufio using a writer buffer
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	str := "Hello, Golang! dhdtfrhdetdfbgdfgdf"

// 	n, err := writer.WriteString(str)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Flush the buffer to ensure data is written to the file
// 	writer.Flush()

// 	fmt.Printf("Wrote %d bytes to file.\n", n)
// }

// WriteString method

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	str := "Hello, Golang!"
// 	n, err := file.WriteString(str)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Wrote %d bytes to file.\n", n)
// }

//// --------------diff read methods
// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     file, err := os.Open("example.txt")
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

//     buf := make([]byte, 100)
//     n, err := file.Read(buf)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Printf("Read %d bytes from file: %s\n", n, buf[:n])
// }

////io.readfile

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("example.txt")

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	data, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Read %d bytes from file: %s\n", len(data), data)
// }

// using bufio

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Read line:", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
