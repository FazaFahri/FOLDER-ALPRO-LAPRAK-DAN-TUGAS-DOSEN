package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%7 == 0 {
		mobil := n /7
		fmt.Print(mobil)
	}else {
		mobil := (n/7) + 1
        fmt.Print(mobil)
	}

}