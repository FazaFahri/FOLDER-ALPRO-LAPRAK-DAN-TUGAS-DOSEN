package main

import "fmt"

func main() {
	var kar rune
	fmt.Scanf("%c", &kar)
	if  ('a' <= kar && kar <= 'z') || ('A' <= kar && kar <= 'Z') {
		fmt.Print("Alphabet")
	}else {
		fmt.Print("Bukan alphabet")
	}


}