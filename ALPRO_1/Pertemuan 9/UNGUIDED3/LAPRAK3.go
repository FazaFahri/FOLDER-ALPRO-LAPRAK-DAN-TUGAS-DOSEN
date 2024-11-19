package main

import "fmt"

func main() {
	var x, y int
	var hasilX, hasilY bool
	fmt.Scan(&x,&y)
	
    hasilX = y%x == 0
	hasilY = x%y == 0

	fmt.Println(hasilX)
	fmt.Println(hasilY)
	
}