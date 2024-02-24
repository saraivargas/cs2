package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Simple Calculator!")

	for {
		selectOperation()

		userOperation := getUserInput()

		if userOperation < 1 || userOperation > 4 {
			fmt.Println("Please try again")
			continue
		}

		fmt.Println("Your choice:", userOperation)

		fmt.Print("Enter the first number: ")
		firstNumber := getUserInput()

		fmt.Print("Enter the second number: ")
		secondNumber := getUserInput()

		fmt.Print("Result: ")

		switch userOperation {
		case 1:
			fmt.Println(firstNumber + secondNumber)
		case 2:
			fmt.Println(firstNumber - secondNumber)
		case 3:
			fmt.Println(firstNumber * secondNumber)
		case 4:
			var floatNum = float32(firstNumber) / float32(secondNumber)
			fmt.Println(floatNum)
		}

	}

}

func selectOperation() {
	fmt.Println("Select an operation:\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division")
}

func getUserInput() int {
	var input int
	fmt.Scan(&input)
	return input

}
