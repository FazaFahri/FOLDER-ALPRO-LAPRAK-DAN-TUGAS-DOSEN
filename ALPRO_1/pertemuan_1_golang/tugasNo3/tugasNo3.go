// variable fxy,x,y type data float 64
// input x,y
// rumus fxy = 1/3x^2+10+10y+7
// output fxy
package main

import (
	"fmt"
	"math"
)

func main() {
	var fxy, x, y float64
	fmt.Scan(&x, &y)
	// fxy = 1/3*math.Pow(x, 2) + 10 + 10*y + 7
	fxy = 1.0/(3.0*math.Pow(x, 2)+10) + 10*y + 7

	fmt.Println("Berikut adalah hasilnya", fxy)
}
