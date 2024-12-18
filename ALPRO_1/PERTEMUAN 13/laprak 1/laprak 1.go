package main

import "fmt"

func main() {
    var angka, jumlah int
    fmt.Println("MASUSKKAN BILANGAN:")
    fmt.Scan(&angka)
    for angka > 0 {
        angka /= 10
        jumlah++
    }
    fmt.Printf("JUMLAH DIGIT %d", jumlah)

}
