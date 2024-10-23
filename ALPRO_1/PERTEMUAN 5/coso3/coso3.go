package main

import "fmt"

func main() {
	var v1, v2 , hasil, j int
	fmt.Scan(&v1,&v2)
	for j = 1; j <= v2; j++ {
		hasil = hasil + v1
	}
	fmt.Print(hasil)
}