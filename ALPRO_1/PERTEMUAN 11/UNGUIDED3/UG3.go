package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch  {
	case n%2 != 0 && n < 10:
		fmt.Println("Kategori bilangan ganjil")
		n2 := n + 1
		hasil1 := n + n2
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", n, "+", n2, "=", hasil1)
	case n%2 == 0 && n < 10:
		fmt.Println("Kategori bilangan genap")
		n2 := n + 1
		hasil1 := n * n2
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", n, "*", n2, "=", hasil1)
	case n%5 == 0 && n> 10 && n%2 != 0:
		fmt.Println("Kategori bilangan kelipatan 5")
		hasil2 := n * n
		fmt.Println("kelipatan 5, hasil kuadrat dari", n, "^2", hasil2)
    case n %10 == 0 && n>10 && n%2==0 :
        fmt.Println("Kategori bilangan kelipatan 10")
		n2 := n / 2
		hasil2 := n / n2
		fmt.Println("hasil pembagian antara ", n, "/", n2,"=", hasil2)

	}


}
