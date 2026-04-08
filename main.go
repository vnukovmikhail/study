package main

import (
	"fmt"
)

func main() {
	message := "Hello, "

	greetingFn := func(name string) {
		fmt.Println(message + name)
	}

	defer greetingFn("Alice")
	defer greetingFn("Bob")

	// os.Exit(1)

	defer func(name string) {
		fmt.Println(message + name)
	}("Lesh")

	fmt.Println("Test")

	message = "Hi, "
}
