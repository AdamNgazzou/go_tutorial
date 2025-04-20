package main

import (
	"errors"
	"fmt"
)

func main() {
	PrintMe("hey")
	var result, remainder, err = intDivision(10, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("the result of the integer division is %v", result)
	} else {
		fmt.Printf("the result of the integer division is %v with the remainder %v", result, remainder)
	}
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a int, b int) (int, int, error) {
	var err error

	if b == 0 {
		err = errors.New("cannot divide by zero ")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
