package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy, ab, bc, ca float64

	fmt.Println("Masukan Format ax dan ay")
	fmt.Scan(&ax, &ay)

	fmt.Println("Masukan format bx dan by")
	fmt.Scan(&bx, &by)

	fmt.Println("Masukan format cx dan cy")
	fmt.Scan(&cx, &cy)

	ab = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc = math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca = math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	// fmt.Println(ab)
	// fmt.Println(bc)
	// fmt.Println(ca)

	fmt.Printf("panjang sisi terpanjang  %.0f", math.Max(math.Max(ab, bc), ca))

}
