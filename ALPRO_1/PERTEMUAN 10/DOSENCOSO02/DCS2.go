package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5 float64
	fmt.Scan(&n1, &n2, &n3, &n4, &n5)
	if n1 < n2 && n2 < n3 && n4 < n5 {
		fmt.Println("Stabil naik")
	} else if n1 > n2 && n2 > n3 && n4 > n5 {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}
