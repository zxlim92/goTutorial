package main

import (
	"fmt"
	"go_tutorials/shared"
	"go_tutorials/tut1"
	// Import the shared package
)

const pi float32 = 3.14

/* Define an interface called Shape with one method:*/

type Shape interface {
	getArea() float64
}

/* Create two structs:

Rectangle with fields Width and Height

Circle with field Radius */

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) getArea() float64 {
	return float64(r.Width) * float64(r.Height)
}

type Circle struct {
	Radius int
}

func (c Circle) getArea() float64 {
	radius := float64(c.Radius)
	return radius * radius * float64(pi)
}

type Square struct {
	Rectangle Rectangle
}

func NewSquare(l int) (Square, error) {
	if l <= 0 {
		return Square{}, fmt.Errorf("side must be positive")
	}
	r := Rectangle{l, l}
	return Square{r}, nil
}

func (s Square) getArea() float64 {
	return s.Rectangle.getArea()
}

//Make both structs implement the Shape interface.

func main() {

	r := Rectangle{10, 7}

	var c Circle
	c = Circle{5}

	slice := []Shape{}
	square, err := NewSquare(3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	slice = append(slice, c)
	slice = append(slice, r)
	slice = append(slice, square)
	for i, s := range slice {
		fmt.Println(i, s.getArea())
	}

	// Using imported public method from shared package
	fmt.Println(shared.PublicMethod())
	fmt.Println(tut1.PublicMethod())
}
