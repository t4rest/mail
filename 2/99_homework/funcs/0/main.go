package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	z:=1.0
	min_delta := 0.00000000001
	delta := 1.0
	i := 0
	for ; math.Abs(delta) > min_delta ; i++{
		delta = (z*z - x)/(2*z)
		//fmt.Println(z,delta)
		z = z - delta
	}
	fmt.Println("iterations: ", i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
