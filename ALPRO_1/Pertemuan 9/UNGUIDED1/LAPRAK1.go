package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		motor := n / 2
		fmt.Print(motor)
	} else {
		motor := (n / 2) + 1
		fmt.Print(motor)
	}
}
