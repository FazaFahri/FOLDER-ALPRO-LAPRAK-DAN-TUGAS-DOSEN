package main

import "fmt"

type barang struct {
	namaBarang  string
	jumlah      int
	hargaSatuan float64
	totalHarga  float64
}

func main() {
	var b barang
	fmt.Println ("nama barang")
	fmt.Scan(&b.namaBarang)

	fmt.Println("yang di ambil")
	fmt.Scan(&b.jumlah)

	fmt.Println("harga satuan")
    fmt.Scan(&b.hargaSatuan)

 b.totalHarga = float64(b.jumlah) * b.hargaSatuan

 fmt.Printf("total nya adalah %.2f" ,b.totalHarga)

}