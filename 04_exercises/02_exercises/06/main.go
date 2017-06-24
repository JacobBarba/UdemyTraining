package main

import (
	"fmt"
)

func even(a int) bool {
	return a%2 == 0
}

func less(a *int, b *int) *int {
	if *a < *b {
		return a
	} else {
		return b
	}
}

func main() {
	a := 1
	b := 2
	answer := 2

	for a+b < 4000000 {
		if even(a + b) {
			answer += a + b
		}
		*less(&a, &b) = a + b
	}
	fmt.Println(answer)
}

//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
