package main

import "fmt"

func main() {
	var intPtr *int
	fmt.Println(intPtr)

	age := 10
	intPtr = &age
	fmt.Println(intPtr)
	fmt.Println(*intPtr)
}
