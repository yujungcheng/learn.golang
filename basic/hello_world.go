package main

import (
	"fmt"
)

var s = "Hello World2!"

func main() {
	fmt.Println("Hello World!")

	i := 42
	fmt.Println(s)
	fmt.Println(i)

	fmt.Printf("Printing 's' variable from outer block %v\n", s)
	b := true
	if b {
		fmt.Printf("Printing 'b' variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Printing 'i' variable from outer block %v\n", i)
			fmt.Printf("Printing 'b' variable from outer block %v\n", b)
			fmt.Printf("Printing 's' variable from outer block %v\n", s)
		}
	}

	fmt.Println(&s)
}
