package main

import (
	"fmt"
	"math"
)

type square struct {
	height float64
	lenght float64
}

type circle struct {
	radius float64
}

type gShape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * (c.radius * c.radius)

}

func (s square) area() float64 {
	return s.lenght * s.height

}

func info(s gShape) {
	fmt.Println(s.area())
}

func main() {
	sqr1 := square{
		height: 4,
		lenght: 2,
	}

	crcl1 := circle{
		radius: 5,
	}

	info(sqr1)
	info(&crcl1)

}

// when I have pointers receptors we can only use pointer type values (*int)
// when I have non-pointer receptors I can use both non-pointer and pointer values (int or *int)
