package main

import "fmt"

func main() {
	var kendaraan string
	var jam int
	fmt.Println("Jenis kendaraan")
	fmt.Scanf("%s",&kendaraan)
    
	fmt.Println("Masulkan durasi")
	fmt.Scan(&jam)
	switch {
	case kendaraan == "motor" :
		hasil := jam * 2000
		fmt.Println("Rp.", hasil)
	case kendaraan == "mobil" :
		hasil := jam * 5000
		fmt.Println("Rp.", hasil)
	case kendaraan == "truk" :
		hasil := jam * 8000
		fmt.Println("Rp.", hasil)
	}
}