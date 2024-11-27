package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)
	if n1 == n2 && n2 == n3 && n3 == n1 {
		fmt.Println("Segitiga Sama Sisi")
	} else if n1 != n2 && n2 != n3 && n3 == n1 {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}
