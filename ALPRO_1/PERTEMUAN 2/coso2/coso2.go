package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Scan(&alas, &tinggi)
	luas = (alas * tinggi / 2)
	fmt.Println("Maka luas nya adalah : ",luas)
}