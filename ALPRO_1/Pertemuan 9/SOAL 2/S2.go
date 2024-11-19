package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	if bilangan == 100 {
		fmt.Print("benar")
	} else {
		fmt.Print("salah")
	}
	
}