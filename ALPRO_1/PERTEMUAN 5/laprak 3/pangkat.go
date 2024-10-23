package main

import "fmt"

func main() {
	var v1, v2, hasil int
	fmt.Scan(&v1, &v2)
	hasil = 1
	for i := 1; i <= v2; i++ {
		hasil *= v1

	}
	fmt.Print(hasil)
}
