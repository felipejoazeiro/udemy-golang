package main

import (
	"fmt",
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma){
	fmt.Printf("A área da forma é %0.2\n", f.area())
}