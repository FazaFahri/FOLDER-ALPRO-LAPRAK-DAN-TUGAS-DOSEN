package main

import "fmt"

func main() {
	var token string
	fmt.Scan(&token)
	for token != "12345abcde" {
		fmt.Println("sandi yang di masukan salah")
		fmt.Scan(&token)
	}
	fmt.Print("selamat anda berhasil login")
}