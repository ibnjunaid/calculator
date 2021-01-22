package funcs

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func division(a, b int) int {
	if b == 0 {
		err := fmt.Errorf("Divide by zero Error")
		panic(err)
	}
	return a / b
}

func modulo(a, b int) int {
	if b == 0 {
		err := fmt.Errorf("Modulo by zero Error")
		panic(err)
	}
	return a % b
}
