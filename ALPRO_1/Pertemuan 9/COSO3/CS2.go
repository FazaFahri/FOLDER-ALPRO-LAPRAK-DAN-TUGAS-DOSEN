package main

import "fmt"

func main() {
	var kar rune
	fmt.Scanf("%c",&kar)
	
	if (kar != 'a' && kar !='i' && kar !='u' && kar !='e' && kar !='o') && (kar != 'A' && kar != 'I' && kar != 'U' && kar !='E' && kar != 'O' ) {
		fmt.Print("konsonan")
	}else {
		fmt.Print("bukan konsonan")
	}
}

// KONSONAN