package main

import (
	"fmt"
)

func main() {
	var celcius, fahrenheit, reamur, kelvin float64
	fmt.Println("Masukan celcius")
	fmt.Scan(&celcius)
	fmt.Println("Masukan fahrenheit")
	fmt.Scan(&fahrenheit)

	celcius = (fahrenheit-32)* 5/9
	reamur = celcius * 4/5
	kelvin = (fahrenheit - 32)*5/9 +273.15 
	fmt.Println("hasil dari celcius", celcius)
	fmt.Println("Hasil dari reamur", reamur)
    fmt.Println("Hasil dari fahrenheit", fahrenheit)
	fmt.Println("Hasil dari kelvin", kelvin)
}

// selesai