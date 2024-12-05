package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	switch{
	case n>= 6.5 && n <= 8.6 :
		fmt.Println("Air Layak Minum")
	case n >= 9 && n <= 14 :
		fmt.Println("Air Tidak Layak Minum")
	default :
		fmt.Println("Nilai pH Tidak Valid , Nilai pH harus antara 0 sampai 14")
	}
}