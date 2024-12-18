package main

import "fmt"

func main() {
	var x, y, hitung int
	fmt.Scan(&x,&y)
	hitung = 0
	for x >= y {
		x -= y
		hitung += 1
	}
	fmt.Println(hitung)
}