package main

import "fmt"


func oddeven(x int) {
	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}

func main() {
	x := 12 % 3

	oddeven(x)

	y := 13 % 3

	oddeven(y)
}
