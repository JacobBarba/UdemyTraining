package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&a)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&b)

	var c int
	c = b % a
	fmt.Printf("The remainder is: %d\n", c)
}
