package main

import (
	"fmt"
)

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(half(i))
	}
}
