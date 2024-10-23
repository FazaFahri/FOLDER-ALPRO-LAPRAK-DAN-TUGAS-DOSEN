package main

import "fmt"

func main() {
	var belanjaAwal, diskon, besarDiskon, hargaAkhir int
	fmt.Println("Masukan Harga Belanja")
	fmt.Scan(&belanjaAwal)

	fmt.Println("Masukan diskon")
	fmt.Scan(&diskon)

	besarDiskon = belanjaAwal *  diskon / 100
	hargaAkhir = belanjaAwal - besarDiskon 

	fmt.Println("hasil nya adalah = ", hargaAkhir)
	
}