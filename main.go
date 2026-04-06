package main

import (
	"fmt"
)

// var myInt2 int8 = 10

var (
	myInt2    int    = 10_000_000
	myString2 string = "Hello World!"
)

func main() {
	var myString string
	fmt.Println(myString)
	myString = "welcome to the code"
	fmt.Println((myString))

	var smalPositiveValue uint8
	smalPositiveValue = 255
	fmt.Println(smalPositiveValue)

	var smalNegativeValue int8
	smalNegativeValue = -128
	fmt.Println(smalNegativeValue)

	var myInt int
	fmt.Println(myInt)
	myInt = int(smalNegativeValue)
	myInt = int(smalPositiveValue)

	var myByte byte = 'A'
	fmt.Println(myByte)

	var smallFloat float32
	fmt.Println(smallFloat)
	smallFloat = 13.231921
	fmt.Println(smallFloat)

	var bigFloat float64
	fmt.Println(bigFloat)
	bigFloat = 43.49348394834933333333
	fmt.Println(bigFloat)

	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPart)

	myString = `Hello
	
	World!`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "learn"
	lastName = "code"

	var fullName string
	// fullName = firstName + " " + lastName

	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)

	fmt.Println(myInt2, myString2)
	myInt2 := 15
	fmt.Println(myInt2)

	a := [...]int{3, 5: 10, 9, 1}
	fmt.Println(a, a[3], a[:1])

	var x []int
	fmt.Println(x == nil)

	x = []int{}
	fmt.Println(x == nil)

	y := []int{1, 2: 10, 0, 3}
	y = []int{5, 3: 7, 4, 1}
	fmt.Println(y)

	z := make([]int, 5)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	t := make([]string, 5, 10)
	fmt.Println(t)

	r := make([]int, 5)
	fmt.Println(r)
	r = append(r, 3)
	r = append(r, 3, 5, 7, 8)
	fmt.Println(r)
	fmt.Println(cap(r))

	b := []int{3, 4, 5}
	r = append(r, b...)
	fmt.Println(r)
	r[5] = 1

	fmt.Println(r[:6])
	fmt.Println(b[2:])

	e := make([]int, 6)
	copy(e, r)
	fmt.Println(e)

	// var nameAge map[string]int
	// nameAge["foo"] = 21
	// fmt.Println(nameAge["foo"])

	// nameAge := map[string]int{}
	var nameAge map[string]int = map[string]int{
		"a": 1,
		"b": 6,
		"c": 9,
	}
	nameAge["foo"] = 21
	nameAge["bar"] = 33
	nameAge["foo bar"] = 17
	fmt.Println(len(nameAge))
	fmt.Println(nameAge["c"])

	gradeC, ok := nameAge["c"]
	fmt.Println(gradeC, ok)

	s := map[string][]int{
		"foo": []int{1, 2, 3, 4},
	}
	fmt.Println(s)
}
