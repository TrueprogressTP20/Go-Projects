package main
import (
	"fmt";"os";"strconv"
)

func main() {
	// If wrong format in the terminal
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		return
	}

	// strconv package convert the command-line arguments("os.Args") to float-point numbers(`float64`)
	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)
	operator := os.Args[2]

	// Checking if both operands are number
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers provided")
		return
	}

	// Initializing result
	result := 0.0

	// Using Switch for better option making
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
				fmt.Println("Error: Division by zero is not allowed.")
				return
			}
		default:
			fmt.Println("Invalid operator. Please use +, -, *, or /.")
			return
	}

	// Output out result
	fmt.Printf("%.1f %s %.1f = %.1f\n", num1, operator, num2, result)

}