package main

import "fmt"

func main() {
	var n int
	var hasil bool
	fmt.Scan(&n)
	hasil = false
	if n < 0 && n%2 == 0 {
		hasil = true
	}
	fmt.Print(hasil)

}