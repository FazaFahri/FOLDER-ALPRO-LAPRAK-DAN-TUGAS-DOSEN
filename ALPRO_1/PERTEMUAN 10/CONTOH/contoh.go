package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a,&b)
	if b != 0 {
		hasil := a / b
		fmt.Println(hasil)
	}else {
		fmt.Println("b Bernilai nol")
	}
	fmt.Println("program selesai")
}