package main

import (
	"fmt"
)

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(half(i))
	}
}
