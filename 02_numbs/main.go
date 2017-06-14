package main

import "fmt"

func main(){
	for i := 0; i <200; i++ {
		fmt.Printf("%s %d %b %x \n", "goat", i, i, i)
	}
}