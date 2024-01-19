package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"a", "b", "c", "d"}
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "e", "f")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3], "g")
	fmt.Println(fruitlist)

	//updating value
	fruitlist[0] = "z"
	fmt.Println(fruitlist)
	// fruitlist = append(fruitlist, )

	//delete an elem
	fruitlist = append(fruitlist[:1], fruitlist[1+1:]...)
	fmt.Println(fruitlist)

	// create slices using make
	slice := make([]int, 4)

	slice[0] = 1
	slice[2] = 11
	slice[3] = 111
	slice[1] = 1111
	fmt.Println(slice)
	//slice[5] = 11111 this will give you error because len is 4

	//but instead of this we can use append method to add elem withou err
	slice = append(slice, 55, 66)
	fmt.Println(slice)

	//sort a slice

	sort.Ints(slice)
	fmt.Println(slice)
	fmt.Println(sort.IntsAreSorted(slice))

	fmt.Println(slice[:])
	fmt.Println(slice[:5])
	fmt.Println(slice[0:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[0:])

	//deleting elems in slices

	slice = append(slice[:5], slice[6:]...)
	fmt.Println("slice", slice)

	// 2-D slices

	row_1 := []string{"man", "patel", "20", "web dev"}
	row_2 := []string{"jemish", "khunt", "21", "android dev"}
	row_3 := []string{"nil", "patel", "24", "doctor"}

	table := [][]string{row_1, row_2, row_3}

	fmt.Println(row_1)
	fmt.Println(row_2)
	fmt.Println(row_3)
	fmt.Println(table)

	// assigning value of one slice to other

	temp1 := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	temp2 := temp1

	fmt.Println(temp1) //1.2.3.4.5.6.7.8.9
	fmt.Println(temp2) //1,2,3,4,5,6,7,8,9

	temp2[0] = 0
	fmt.Println(temp1) // [0] ,2,3,4,5,6,7,8,9
	fmt.Println(temp2) // [0] ,2,3,4,5,6,7,8,9

	//to only change in second slice use "  copy  " method

	a1 := []int{7, 8, 9}
	a2 := make([]int, 3)

	copy(a2, a1) /////

	a2[0] = 0

	fmt.Println("slicwe 1", a1)
	fmt.Println("slice 2", a2)
}
