package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for loop:= false; !loop; {
		fmt.Scan(&n)
		loop = n >= 0
	}
	fmt.Printf("%d adalah bilangan positif\n", n)
}