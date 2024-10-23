// Variable r,luaslingkaran,kelilinglingkaran tipe data float 64
// const phi dengan value 3.14
// input r
// rumus luaslingkaran = phi*r^2
// rumus kelilinglingkaran = 2 * phi * r
// output luaslingkaran, kelilinglingkaran



package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luaslingkaran, kelilinglingkaran float64
	const phi = 3.14

	fmt.Println("Masukan jari-jari")
	fmt.Scan(&r)

	luaslingkaran = phi * math.Pow(r, 2)
	kelilinglingkaran = 2 * phi * r

	fmt.Println("hasil dari luas lingkaran adalah", luaslingkaran)
	fmt.Println("hasil dari keliling lingkaran adalah", kelilinglingkaran)
}
