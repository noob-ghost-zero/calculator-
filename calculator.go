package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Ask the user for input
	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter an operator (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	// Perform the calculation
	var result float64
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			err = fmt.Errorf("cannot divide by zero")
		}
	default:
		err = fmt.Errorf("invalid operator")
	}

	// Print the result or error
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
