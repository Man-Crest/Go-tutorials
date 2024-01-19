package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "saturday", "sunday"}
	fmt.Println(days)

	//for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//similer to for in loop

	for j := range days {
		//returns index j not value
		fmt.Println(days[j])
	}

	//similer to froeach
	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	//while
	initialvalue := 1

	for initialvalue <= 10 {

		if initialvalue == 5 {
			goto temp
		}

		fmt.Println("the value is", initialvalue)
		initialvalue++
	}

	// this is lable we can name it anything and used for goto statements

temp:
	fmt.Println("this is written inside the lable")
}
