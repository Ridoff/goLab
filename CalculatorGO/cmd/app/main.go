package main

import (
	"fmt"
	"go/course/calcu/internal/calculator"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	result, err := calculator.Calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
  