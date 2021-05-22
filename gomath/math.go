package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func sub(number1 int, number2 int) int {
	return number1 - number2
}

func times(number1 int, number2 int) int {
	return number1 * number2
}

func div(number1 int, number2 int) int {
	return number1 / number2
}
