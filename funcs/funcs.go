package funcs

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Product(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	if b == 0 {
		err := fmt.Errorf("Divide by zero Error")
		panic(err)
	}
	return a / b
}

func Modulo(a, b int) int {
	if b == 0 {
		err := fmt.Errorf("Modulo by zero Error")
		panic(err)
	}
	return a % b
}
