package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	height float64
	width  float64
}

type triangle struct {
	side float64
}

func (t triangle) area() float64 {
	fmt.Println("area of triangle ", 3*t.side)
	return 3 * t.side
}
func (t triangle) circumference() float64 {
	fmt.Println("circumference of triangle ", 3*t.side)
	return 3 * t.side
}

func (c circle) area() float64 {
	fmt.Println("area of circle are ", 3.14*c.radius*c.radius)
	return 3.14 * c.radius * c.radius
}
func (c circle) circumference() float64 {
	fmt.Println("circumference of circle are ", 2*3.14*c.radius)
	return 2 * 3.14 * c.radius
}

func (s square) area() float64 {
	fmt.Println("area of square are ", s.height*s.width)
	return s.height * s.width
}
func (s square) circumference() float64 {
	fmt.Println("circumference of square are ", 2*s.width)
	return 2 * s.width
}

type shape interface {
	area() float64
	circumference() float64
}

func main() {
	c := circle{4}
	s := square{2, 2}
	t := triangle{3}

	var interfaces shape = c
	interfaces.area()
	slices := []shape{c, s, t}

	slices[0].area()
	slices[1].area()

	for _, shape := range slices {
		shape.area()

		// val, ok := shape.(square)
		// if ok {
		// 	val.circumference()
		// }
	}
	describe(c)
	describe(s)
}

// an empty interface without any methods can return or take any type as input

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
