package main

import "fmt"

func main() {
	var n int
	var tof bool
	fmt.Scan(&n, &tof)
	if tof == true && n > 100000 && n >= 200000 {
		diskon := (n * 90) / 100
		hasil := diskon - 75000

		fmt.Println("kartu", tof)
		fmt.Println("diskon", n > 100000)
		fmt.Println("cashback", tof)
		fmt.Println(hasil)
	} else {
		hasil := (n * 90) / 100
		fmt.Println("kartu", tof)
		fmt.Println("diskon", n > 100000)
		fmt.Println("cashback", tof)
		fmt.Println(hasil)
	}
}
