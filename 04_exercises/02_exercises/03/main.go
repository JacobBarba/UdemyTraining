package main

import (
	"fmt"
)

func greatest(a ...int) int {
	var b int
	for _, v := range a {
		if v > b {
			b = v
		}
	}
	return b
}

func main() {
	fmt.Println(greatest(1, 9, 3, 4, 5, 6))
}
