package main

import "fmt"

// This is a global variable since it's declared outside a function
var number int = 5

func main() {
	var decision bool = true // a local scope variable since it's declare inside a function
	fmt.Printf("original value of number:%d\n", number)
	number = 10 //re-assigning value of number to local scope variable
	fmt.Printf("new value of number: %d\n", number)
	fmt.Printf("value of decision: %t\n", decision)
}
