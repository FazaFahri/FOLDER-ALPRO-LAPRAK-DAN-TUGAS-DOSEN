package main

import "fmt"

func main() {
	var tinggi int
	fmt.Scan(&tinggi)

	i := tinggi

	for i >= 1 {
        j := 1
		for j <= i {
			fmt.Print("*")
			j++
		}
		fmt.Println()
		i--
	}
	
}