package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	if n < 2.00 {
		fmt.Println("Poor")
	} else if n >= 2.00 && n <= 2.75 {
		fmt.Println("Fair")
	} else if n >= 2.76 && n <= 3.00 {
		fmt.Println("Satisfactory")
	} else if n >= 3.01 && n <= 3.50 {
		fmt.Println("Very good")
	} else {
		fmt.Println("Excellent")
	} 
	
}
