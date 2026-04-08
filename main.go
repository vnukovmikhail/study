package main

import (
	"errors"
	"fmt"
)

type student struct {
	firstName string
	lastName  string
}

func main() {
	a := 10
	increment(&a)
	fmt.Println(a)

	s := student{
		firstName: "code",
		lastName:  "learn",
	}

	previousLastName, err := updateLastName(&s, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(previousLastName)
	// fmt.Println(s)

	prettyPrintStudent(&s)
}

func prettyPrintStudent(s *student) {
	fmt.Printf("F-%s-%s\n", s.firstName, s.lastName)
}

func updateLastName(s *student, newLastName string) (*string, error) {
	if newLastName == "" {
		return nil, errors.New("Empty last name")
	}
	previous := s.lastName
	s.lastName = newLastName
	return &previous, nil
}

func increment(x *int) {
	*x++
}
