package main

import "fmt"

func main() {
	var full, half int
	var predikat string
	fmt.Scan(&full)

	switch  {
	case full == 0:
		half = 12 
		predikat = "AM"
	case full < 12 :
		half = full
	     predikat = "AM"
	case full == 12 :
		half = 12
		predikat = "PM"
	case full > 12 :
		half = full - 12
		predikat = "PM"
		
	}
	fmt.Println(half,predikat)
}