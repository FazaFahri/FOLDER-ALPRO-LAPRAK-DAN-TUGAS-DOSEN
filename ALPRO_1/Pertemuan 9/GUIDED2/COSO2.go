package main

import "fmt"

func main() {
	var n int
	var text string
	fmt.Scan(&n)
	text = "negatif"
	if n >= 0 {
		text = "Positif"

	}
	fmt.Print(text)
}

// if n > 0 {
// 	fmt.Print("Positif")
// }else {
// 	fmt.Print("negatif")
// }
