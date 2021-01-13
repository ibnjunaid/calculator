package main

import (
	"fmt"
)

func main() {
	calculateUsing3Operands()
	//calculateUsing3OperandsAndSameOp()
	// calculateUsingSwitch()
	//calculateUsingIfelse()

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
	} else if op == "-" {
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

func calculateUsing3OperandsAndSameOp() {
	var a, b, c int
	var op1, op2 string
	fmt.Print("enter three numbers\n")
	fmt.Scanf("%d %s %d %s %d", &a, &op1, &b, &op2, &c)
	switch op1 {
	case "+":
		fmt.Println(a + b + c)
	case "-":
		fmt.Println(a - b - c)
	case "*":
		fmt.Println(a * b * c)
	case "/":
		fmt.Println(a / b / c)
	case "%":
		fmt.Println(a % b % c)
	}
}

func calculateUsing3Operands() {
	var a, b, c int
	var op1, op2 string
	fmt.Println("enter three numbers")
	fmt.Scanf("%d %s %d %s %d", &a, &op1, &b, &op2, &c)
	if op1 == "+" && op2 == "+" {
		fmt.Println(a + b + c)
	} else if op1 == "+" && op2 == "-" {
		fmt.Println(a + b - c)
	} else if op1 == "+" && op2 == "*" {
		fmt.Println(a + b*c)
	} else if op1 == "+" && op2 == "/" {
		fmt.Println(a + b/c)
	} else if op1 == "-" && op2 == "+" {
		fmt.Println(a - b + c)
	} else if op1 == "-" && op2 == "-" {
		fmt.Println(a - b - c)
	} else if op1 == "-" && op2 == "*" {
		fmt.Println(a - b*c)
	} else if op1 == "-" && op2 == "/" {
		fmt.Println(a - b/c)
	} else if op1 == "*" && op2 == "+" {
		fmt.Println(a*b + c)
	} else if op1 == "*" && op2 == "-" {
		fmt.Println(a*b - c)
	} else if op1 == "*" && op2 == "*" {
		fmt.Println(a * b * c)
	} else if op1 == "*" && op2 == "/" {
		fmt.Println(a * b / c)
	} else if op1 == "/" && op2 == "+" {
		fmt.Println(a/b + c)
	} else if op1 == "/" && op2 == "-" {
		fmt.Println(a/b - c)
	} else if op1 == "/" && op2 == "*" {
		fmt.Println(a / b * c)
	} else if op1 == "/" && op2 == "/" {
		fmt.Println(a / b / c)
	} else {
		fmt.Println("not found")
	}
}
