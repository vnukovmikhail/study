package main

import (
	"fmt"
)

func main() {
	day := "Tuef"

	switch day {
	case "Mon":
		fmt.Println("Its monday")
	case "Tue":
		fmt.Println("Oh my ga")
	default:
		fmt.Println("Idonow")
	}

	switch wordLen := len(day); wordLen {
	case 4:
		fallthrough
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(22)
	case 3:
		fmt.Println(333)
	}
	/* for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	} */
	/* outerLoop:
	   	for i := 0; i < 3; i++ {
	   		for j := 0; j < 3; j++ {
	   			if i == 1 && j == 1 {
	   				continue outerLoop
	   			}
	   			fmt.Println(i, j)
	   			if i == 2 && j == 2 {
	   				goto end
	   			}
	   			fmt.Println("watafa")
	   		}
	   	}

	   end:
	   	fmt.Println("The end of the program") */
}
