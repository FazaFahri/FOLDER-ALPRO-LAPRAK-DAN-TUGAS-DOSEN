package main

import "fmt"

type nilai struct {
	panjang  float64
	lebar    float64
	luas     float64
	keliling float64
}

func main() {
	var t nilai
	fmt.Println("Masukan panjang")
	fmt.Scan(&t.panjang)
	fmt.Println("Masukan lebar")
	fmt.Scan(&t.lebar)

	t.luas = t.panjang * t.lebar
	t.keliling = 2 * (t.panjang + t.lebar)

	fmt.Printf("hasil luas adalah %.1f", t.luas)
	fmt.Printf("hasil dari keliling adalah %.1f", t.keliling)
}
