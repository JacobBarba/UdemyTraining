package main

import (
	"fmt"
)

func main() {
	ag := 0
	for i := 0; i < 1000; i++ {
		switch {
		case i%3 == 0 || i%5 == 0:
			ag += i
		}
	}

	fmt.Printf("%d\n", ag)
}
