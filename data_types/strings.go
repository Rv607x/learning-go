package main

import "fmt"

func main() {
	var cow string = "Hello World!"
	fmt.Println(cow)
	fmt.Println(len(cow)) //lenth of string

	// string concatenation
	s := "hell" + "o "
	s += "world"
	fmt.Println(s)
	fmt.Println(cow[0])
}
