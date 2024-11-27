package main

import "fmt"

func main() {
	var n int
	var tof bool
	fmt.Scan(&n,&tof)
	if n >= 17 && tof == true{
		fmt.Println("bisa membuat ktp")
	}else{
		fmt.Println("belum bisa membuat ktp")
	}
}