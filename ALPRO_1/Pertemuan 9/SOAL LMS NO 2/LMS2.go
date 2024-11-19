package main

import "fmt"

func main() {
	var n int
	var tof bool
	fmt.Scan(&n)
	fmt.Scan(&tof)
    if tof == true {
		hasil := (n*65) / 100 
		fmt.Print("Sedang assesment ",hasil)
	} else {
		fmt.Print("tidak sedang assesment ",n)
	}
}