package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World from the main function!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Printf("%s", err.Error())
	case remainder == 0:
		fmt.Printf("This is the result of the integer division %v", result)
	default:
		fmt.Printf("This is the result of the integer division %v with the remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("\n\n\nNo remainder")
	case 1, 2:
		fmt.Println("Remainder is 1 or 2")
	default:
		fmt.Println("Remainder is greater than 2")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero. try again")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
