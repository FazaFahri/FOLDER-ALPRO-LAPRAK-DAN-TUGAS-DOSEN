package main

import "fmt"

func main() {
	var user, password string
	n:= 0 
	for loop:= false; !loop; {
		fmt.Scan(&user, &password)
		loop = user == "Admin" && password =="Admin"
		n = n + 1
	}
	fmt.Print(n, " Percobaan login")
}
