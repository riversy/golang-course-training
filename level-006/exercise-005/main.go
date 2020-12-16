package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type square struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	return s.height * s.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(sh shape) {
	area := sh.area()
	fmt.Printf("This shape's area is %.2f meters\n", area)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := circle{
		radius: rand.Float64(),
	}
	info(c)

	sq := square{
		height: rand.Float64(),
		width:  rand.Float64(),
	}
	info(sq)
}
