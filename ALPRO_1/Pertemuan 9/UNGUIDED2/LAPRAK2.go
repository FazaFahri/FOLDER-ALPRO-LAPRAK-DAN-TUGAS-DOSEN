package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if (n%2 == 0) && (n < 0) {
		fmt.Print("Genap Negatif")
	} else {
		fmt.Print("Bukan")
	}
}
