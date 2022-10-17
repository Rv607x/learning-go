package main

import "fmt"

// This is a global variable since it's declared outside a function
var number int = 5

func main() {
	var decision bool = true // a local scope variable since it's declare inside a function
	fmt.Println("original value of number: ", number)
	number = 10 //re-assigning value of number to local scope variable
	fmt.Println("new value of number: ", number)
	fmt.Println("value of decision: ", decision)
}
