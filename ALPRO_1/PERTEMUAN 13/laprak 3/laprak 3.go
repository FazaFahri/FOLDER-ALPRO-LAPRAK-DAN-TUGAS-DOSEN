package main

import "fmt"

func main() {
    var target int
    var sumbangan int
    var total int = 0
    var donatur int = 0

    fmt.Scan(&target)

    for total < target {
        fmt.Scan(&sumbangan)
        donatur++
        total += sumbangan
        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n",
            donatur, sumbangan, total)
    }

    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n",
        total, donatur)
}

