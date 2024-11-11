package main

import (
	"fmt"
	"math/rand/v2"
)

func isEmpty(a any) bool {
	return a == ""
}

func nums(a int) int {
	return rand.IntN(a)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return b / a
}

func main() {

	var answer int
	var command string

	var a int = nums(10)
	var b int = nums(10)

	fmt.Println("Pick [A, S, M, D]")

	fmt.Scanln(&command)

	if isEmpty(command) {
		fmt.Println("Empty..")
	}

	if command == "A" {
		fmt.Println("What is:", a, "+", b)

		fmt.Println("Enter Answer")

		fmt.Scan(&answer)

		if isEmpty(answer) {
			fmt.Println("Empty..")
		}

		if answer == add(a, b) {
			fmt.Println("Correct", a, "+", b, "=", add(a, b))
		}
	}

	if command == "S" {
		fmt.Println("What is:", a, "-", b)

		fmt.Println("Enter Answer")

		fmt.Scan(&answer)

		if isEmpty(answer) {
			fmt.Println("Empty..")
		}

		if answer == sub(a, b) {
			fmt.Println("Correct", a, "-", b, "=", sub(a, b))
		}
	}

	if command == "M" {
		fmt.Println("What is:", a, "*", b)

		fmt.Println("Enter Answer")

		if isEmpty(answer) {
			fmt.Println("Empty..")
		}

		fmt.Scan(&answer)
		if answer == mul(a, b) {
			fmt.Println("Correct", a, "*", b, "=", mul(a, b))
		}
	}

	if command == "D" {
		fmt.Println("What is:", b, "/", a)

		fmt.Println("Enter Answer")

		if isEmpty(answer) {
			fmt.Println("Empty..")
		}

		fmt.Scan(&answer)
		if answer == div(b, a) {
			fmt.Println("Correct", b, "/", a, "=", div(b, a))
		}
	}

}
