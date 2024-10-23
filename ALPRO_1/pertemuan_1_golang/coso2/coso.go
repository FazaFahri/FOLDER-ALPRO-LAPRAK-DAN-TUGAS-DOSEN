package main

import "fmt"

func main() {
	// fx = 2 / (x + 5) + 5
	// masukan input  x

	var x float64
	var fx float64
	
	fmt.Scan(&x)
	fx = 2 / (x + 5) + 5
	
	fmt.Print(fx)

	
}