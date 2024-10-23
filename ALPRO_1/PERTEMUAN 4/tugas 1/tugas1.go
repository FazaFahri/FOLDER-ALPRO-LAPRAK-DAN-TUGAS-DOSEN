package main

import (
	"fmt"

)

type tabung struct {
	tinggi, jari2         int
	luas, volume          float64
	luasalas, luasdinding float64
}

const phi = 3.14159

func main() {
	var t tabung
	fmt.Scan(&t.jari2, &t.tinggi)
	t.luasalas = phi * float64(t.jari2) * float64(t.jari2)
	t.luasdinding = float64(t.tinggi) * (2 * phi * float64(t.jari2))
	t.luas = 2*t.luasalas + t.luasdinding
	t.volume = t.luasalas * float64(t.tinggi)
	fmt.Println(t.luas)
	fmt.Println(t.volume)

}
