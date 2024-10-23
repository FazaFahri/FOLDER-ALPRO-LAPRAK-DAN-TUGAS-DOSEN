package main

import "fmt"

func main() {
	var smt, nilai int
	var status bool

	fmt.Println("Masukan Semester")
	fmt.Scan(&smt)

	fmt.Println("Masukan Nilai")
	fmt.Scan(&nilai)

	status = ((smt < 8) && (nilai > 500))

	fmt.Println("keulusan", status)

	if status {
		fmt.Printf("Mahasiswa lulus dengan kulian selama %d  dan EPrT %d", smt, nilai)
	} else {
		fmt.Printf("Mahasiswa tidak cumalauded karna kuliah hingga %d semester", smt)
	}

}
