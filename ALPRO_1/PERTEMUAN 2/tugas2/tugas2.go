package main

import (
	"fmt"
	"math"
)

func main() {
	var luaskulitbola, volumebola, r float64
	var phi = 3.1415926535
	fmt.Scan(&r)
	luaskulitbola = 4 * phi * math.Pow(r, 2)
	volumebola = (4.00 / 3.00) * phi * math.Pow(r, 3)
	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f",r,volumebola,luaskulitbola)
}