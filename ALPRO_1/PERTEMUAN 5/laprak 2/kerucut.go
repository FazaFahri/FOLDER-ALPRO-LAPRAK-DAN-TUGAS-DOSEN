package main

import (
	"fmt"
	"math"
)
func main() {
	var  tinggi, n ,r int
	var v float64

	
	fmt.Scan(&n)
	
	
	for j := 1; j  <= n ; j ++  {	
		fmt.Scan(&r, &tinggi)
		v =(1.0/3.0) * math.Pi * math.Pow(float64(r),2) * float64(tinggi)
		fmt.Println(v)
	}
}