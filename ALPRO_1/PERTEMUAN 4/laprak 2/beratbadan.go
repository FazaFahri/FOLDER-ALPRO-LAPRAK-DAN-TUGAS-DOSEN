package main

import (
	"fmt"
	"math"
)
func main() {
	var bmi, tinggiBadan, beratBadan float64
	fmt.Println("Masuakn BMI")
	fmt.Scan(&bmi)

	fmt.Println("Masukan tinggi badan")
	fmt.Scan(&tinggiBadan)

	beratBadan = bmi* math.Pow(tinggiBadan,2)
	fmt.Printf("maka berat badan nya adalah :  %.f  kg", beratBadan)
}