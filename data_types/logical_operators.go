package main

import "fmt"

func main() {
	b3 := 10 > 5 //greater than operator
	fmt.Println(b3)
	b3 = 10 < 5 // less than operator
	fmt.Println(b3)
	b3 = 5 <= 10 // less than equal to
	fmt.Println(b3)
	b3 = 10 != 10 // not equal to
	fmt.Println(b3)

	fmt.Println("****************************")
	// Boolean logical operators below
	c3 := 10 > 5 && 7 < 15 // AND Operator
	fmt.Println(c3)
	c3 = 10 < 5 || 2 > 7 // OR Operator
	fmt.Println(c3)
	c3 = !c3 // NOT operator
	println(c3)
}
