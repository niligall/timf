package main

import (
	"fmt"
	"math"
)

func computeFunctionValue(x float64) float64 {
	var result float64

	if x <= -3 {
		result = math.Atan(math.Sqrt(math.Abs(x))) + 3
	} else if x >= -3 && x < 3 {
		result = math.Cos((math.Pow(x, 5) + 2*x) / (3 + x*x))
	} else if x >= 3 {
		result = math.Pow(x, 4) + math.Pow(3, -x)*math.Pow(x, x-5)
	}

	return result
}

func main() {
	var x float64
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)
	y := computeFunctionValue(x)
	fmt.Printf("Значение функции y(x) в точке %.2f: %.4f\n", x, y)
}
