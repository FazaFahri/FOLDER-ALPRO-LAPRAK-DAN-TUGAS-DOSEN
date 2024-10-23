// program mencari luas dan keliling persegi
// var p,l,K,L tipe data int
// input p,l
// rumus L = p * l
//  rumus K = (2 * p) + (2 * l)
// output L,K
// EndProgram



package main

import "fmt"

func main() {
	var p, l, L, K int
	fmt.Scan(&p, &l)
	
	
	K = (2 * p) + (2 * l)
	fmt.Println(K)
	
	L = p * l
	fmt.Println(L)


}
