package main

import "fmt"

func main() {
	var index rune
	var batasA, batasB int
	fmt.Scanf("%c",&index)
	switch index {
	case 'A':
		batasA = 100
		batasB = 75
	case 'B':
		batasA = 75
		batasB = 65
	case 'C':
		batasA = 65
		batasB = 50
	default:
		batasA = 50
		batasB = 0
	}
	fmt.Printf("Rentang nilai %c adalah %v sampai dengan %v", index, batasA, batasB)
}
