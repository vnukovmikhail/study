package main

import "fmt"

func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}

func apply(f func(int) int, x int) int {
	return f(x)
}

func multiplyBy(multiplier int) func(int) int {
	return func(i int) int {
		return i * multiplier
	}
}

func main() {
	multiplyByTwo := multiplyBy(2)
	multiplyByThree := multiplyBy(3)
	fmt.Println(multiplyByTwo(7), multiplyByThree(9))

	sayHi()
	fn := fullname("Mr.", "Bug")
	sayHiToSomeone(fn)

	_, l := fullnameWithLength("Mr.", "Bug")
	fmt.Println(fn, l)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println(s)

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums...))

	fiveTime := func(x int) int {
		return x * 5
	}
	result := apply(fiveTime, 5)
	fmt.Println(result)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func fullnameWithLength(firstName string, lastName string) (string, int) {
	fn := fmt.Sprintf("%s %s", firstName, lastName)
	return fn, len(fn)
}

func sayHi() {
	fmt.Println("Hi!")
}

func sayHiToSomeone(name string) {
	fmt.Println("Hi!", name)
}

func fullname(firstName string, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}
