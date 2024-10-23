package main

import ("fmt"
"math")
func main() {
	var beratbadan, tinggibadan, bmi float64
	fmt.Println("Masukan berat badan anda")
	fmt.Scan(&beratbadan)

	fmt.Println("Masukan tinggi anda")
	fmt.Scan(&tinggibadan)

	bmi = beratbadan / math.Pow(tinggibadan, 2)
	fmt.Printf("BMI anda adalah %.2f", bmi)
}