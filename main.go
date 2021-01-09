package main

import (
	"fmt"
)

func main() {
	calculateUsingIfelse()
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
