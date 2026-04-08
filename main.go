package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

	numbers := []int{1, 2, 3, 4, 5}
	squared := sliceOperation(numbers, func(i int) int {
		return i * i
	})
	fmt.Println(squared)
}

func sliceOperation(s []int, op func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = op(v)
	}
	return result
}
