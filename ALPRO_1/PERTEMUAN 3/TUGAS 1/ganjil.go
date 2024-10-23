package main

import "fmt"

func main() {
	var bilangan int
	var status bool

	fmt.Println("Masukan bilangan")
	fmt.Scan(&bilangan)

	status = bilangan%2 != 0

	fmt.Println("Hasil nya adalah", status)
}