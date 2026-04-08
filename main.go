package main

import "fmt"

type multiplier func(int) int

type operation func(int, int) int

func multiplyBy(m int) multiplier {
	return func(i int) int {
		return i * m
	}
}

func main() {
	multiplyByTwo := multiplyBy(2)
	multiplyByThree := multiplyBy(3)
	fmt.Println(multiplyByTwo(7), multiplyByThree(9))

	var perform operation
	perform = arithmeticOperation("add")
	fmt.Println(perform(5, 7))
}

func arithmeticOperation(op string) func(int, int) int {
	switch op {
	case "add":
		return func(i1, i2 int) int {
			return i1 + i2
		}
	case "subtract":
		return func(i1, i2 int) int {
			return i1 - i2
		}
	default:
		return func(i1, i2 int) int {
			return i1 + i2
		}
	}
}
