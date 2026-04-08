package main

import (
	"fmt"
)

var (
	a = 1
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

	fmt.Println(a)
	{
		a := 12
		fmt.Println(a)
	}
	fmt.Println(a)
	a := "f"
	for index, r_value := range a {
		fmt.Printf("idx:%d, val:%c\n", index, r_value)
	}

	ages := map[string]int{
		"Alice": 30,
		"Bob":   15,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d y.o.\n", name, age)
	}

	b := 100
	for {
		fmt.Println(b)
		b--
		if b == 90 {
			break
		}
	}
}
