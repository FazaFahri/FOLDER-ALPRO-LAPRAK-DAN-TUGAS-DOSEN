package main

import "fmt"

func main() {
    var angka float64
    fmt.Print("Masukkan angka: ")
    fmt.Scanf("%f", &angka)

    mulai := angka - float64(int(angka))

    for i := mulai; i <= 1.0; i += 0.1 {
        fmt.Printf("%.1f\n", i)
    }

    if mulai > 0 {
        fmt.Printf("%.1f\n", angka+1-mulai)
    }
}
