package main

import "fmt"

func main() {
	var angka, cacah int
	fmt.Scan(&angka)
	for angka > 0{
		cacah = angka % 10
		fmt.Println(cacah)
		angka = angka / 10
	}
}