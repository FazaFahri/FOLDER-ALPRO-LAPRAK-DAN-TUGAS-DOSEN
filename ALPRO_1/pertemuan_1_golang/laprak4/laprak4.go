package main

import "fmt"

func main() {
	var C, F float64
	fmt.Println("masukan fahrenheit")
	fmt.Scan(&F)
	C = (F - 32) * 5 / 9
	fmt.Println("hasilnya adalah ", C)

}
