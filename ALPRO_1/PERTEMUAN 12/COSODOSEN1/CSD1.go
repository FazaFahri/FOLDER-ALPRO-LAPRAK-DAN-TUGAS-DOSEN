package main

import "fmt"

func main() {
	var xi, total int
	total = 0
	fmt.Scan(&xi)
	for xi%2 == 0{
		total += xi
		fmt.Scan(&xi)
	}
	fmt.Println(total)
}