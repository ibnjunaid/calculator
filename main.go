package main

import (
	"fmt"
)

func main() {
	calculateUsingSwitch()
	// calculateUsingIfelse()

}

func calculateUsingSwitch() {
	var a, b int
	var op string
	fmt.Print("enter two numbers\n")
	fmt.Scanf("%d %s %d", &a, &op, &b)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "%":
		fmt.Println(a % b)
	default:
		fmt.Println("not found")
	}
}

func calculateUsingIfelse() {
	var a, b int
	var op string
	fmt.Print("enter two numbers\n")
	fmt.Scanf("%d %s %d", &a, &op, &b)
	if op == "+" {
		fmt.Println(a + b)
	} else if op == "+" {
		fmt.Println(a - b)
	} else if op == "*" {
		fmt.Println(a * b)
	} else if op == "/" {
		fmt.Println(a / b)
	} else if op == "%" {
		fmt.Println(a % b)
	} else {
		fmt.Println("not found")
	}
}
