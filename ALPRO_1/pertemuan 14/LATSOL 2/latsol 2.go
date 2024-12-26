package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	if n == 1 {
		fmt.Println("Bukan prima")
		return
	}
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan prima")
	}
}
