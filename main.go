package main

import (
	"fmt"
)

func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		subject   []string
	}

	var student1 student
	student1 = student{"code", "golang", 17, []string{"maths", "science"}}
	fmt.Printf("%+v\n", student1)

	student2 := student{
		firstName: "foo",
		lastName:  "bar",
		age:       19,
	}

	fmt.Println("First name of student2:", student2.firstName)

	student2.subject = append(student2.subject, "arts")
	fmt.Printf("%+v\n", student2)

	guardian := struct {
		firstName string
		lastName  string
	}{
		firstName: "Alex",
		lastName:  "Theo",
	}
	fmt.Printf("%+v\n", guardian)
}
