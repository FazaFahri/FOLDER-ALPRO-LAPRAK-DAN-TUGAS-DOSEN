package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    fmt.Println("Berat Parsel (gram): ", n)
	kg := n / 1000
	g := (n % 1000) / 1
	fmt.Println("Detail Berat", kg, "kg +", g, "g")
	if g >= 500 {
		biayaPengiriman := 10000 * kg
		tambahanBiaya := 5 * g
		DetailBiaya := biayaPengiriman + tambahanBiaya
		fmt.Println("Detail Biaya: Rp.",biayaPengiriman, "+ Rp.",tambahanBiaya )
		fmt.Println("Total Biaya",DetailBiaya)
	}else{
		biayaPengiriman := 10000 * kg
		tambahanBiaya := 15 * g
		DetailBiaya := biayaPengiriman + tambahanBiaya
		fmt.Println("Detail Biaya: Rp.",biayaPengiriman, "+ Rp.",tambahanBiaya )
		fmt.Println("Total Biaya",DetailBiaya)
	}
}