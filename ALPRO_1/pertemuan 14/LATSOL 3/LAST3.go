package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	expected := []string{"merah", "kuning", "hijau", "ungu"}
	scanner := bufio.NewScanner(os.Stdin)
	isSuccess := true

	fmt.Println("Masukkan urutan warna untuk 5 percobaan:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		scanner.Scan()
		input := scanner.Text()
		colors := strings.Split(input, " ")

		if len(colors) != 4 {
			fmt.Println("Input tidak valid. Harus memasukkan 4 warna.")
			isSuccess = false
			break
		}

		for j, color := range colors {
			if color != expected[j] {
				isSuccess = false
				break
			}
		}

		if !isSuccess {
			break
		}
	}
	if isSuccess {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
