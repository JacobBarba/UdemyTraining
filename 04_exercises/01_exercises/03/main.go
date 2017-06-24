package main

import (
	"fmt"
)

func main() {
	var b string
	fmt.Print("Enter your name: ")
	fmt.Scan(&b)
	fmt.Printf("Hello %s!\n", b)
}
