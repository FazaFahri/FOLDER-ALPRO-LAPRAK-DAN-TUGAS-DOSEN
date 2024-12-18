// REPEAT UNTIL
package main

import "fmt"

func main() {
	var n,iterasi int
    iterasi = 0
	fmt.Scan(&n)
	for {
		iterasi++
		fmt.Print(iterasi, " ")
		if iterasi == n {
			break
		}
	}

}



// WHILEDO

// package main

// import "fmt"

// func main() {
// 	var n,iterasi int
// 	iterasi = 0

// 	fmt.Scan(&n)
// 	for iterasi != n {
// 		iterasi = iterasi + 1
// 		fmt.Println(iterasi)
// 	}
// }