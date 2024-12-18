package main

import "fmt"

func main() {
	var word string
	var repetetion int
	fmt.Scan(&word,&repetetion)
	counter := 0 
	for done := false; !done; {
		fmt.Println(word)
		counter++
		done = (counter >= repetetion)
	}
}