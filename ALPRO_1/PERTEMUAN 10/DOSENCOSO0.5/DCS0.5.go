package main

import "fmt"

func main() {
	var val int
	fmt.Scan(&val)
	if val >= 0 {
		fmt.Println("Positive")
	if val %2 == 0 {
		fmt.Print("Genap")
	}else{
		fmt.Print("Ganjil")
	}
} else {

}

}