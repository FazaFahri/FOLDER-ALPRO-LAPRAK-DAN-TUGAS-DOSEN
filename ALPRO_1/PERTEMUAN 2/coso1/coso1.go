package main

import (
	"fmt"
	"math"
)

func main() {
	var sisi, volume float64
	fmt.Scan(&sisi)
     volume = math.Pow(sisi,3)
	// volume = (sisi * sisi * sisi)
	fmt.Println(volume) 
}