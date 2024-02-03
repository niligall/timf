package main

import (
	"fmt"
	"math"
)

func isInUpperRightQuarterRectangle(x, y float64) bool {
	return 0 <= x && x <= 4 && 0 <= y && y <= 5
}

func computeFunctionValue(x float64) float64 {
	firstTerm := math.Sqrt(math.Pow(math.Acos((math.Cos(math.Pow(math.Abs(x), 1.0/6.0)))/(math.Pow(math.Cos(math.Pow(math.Abs(x), 1.0/6.0)), 2)+5)), 3))
	secondTerm := math.Log(math.Abs(math.Pow(3, math.Tan(x)*math.Tan(x))+2)) / math.Log(3)
	result := firstTerm + secondTerm + secondTerm
	return result
}

func main() {
	var x float64

	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)
	y := computeFunctionValue(x)
	inDomain := isInUpperRightQuarterRectangle(x, y)

	fmt.Printf("Значение функции в точке (%.2f): %.4f\n", x, y)
	fmt.Printf("Принадлежит области D: %t\n", inDomain)
}
