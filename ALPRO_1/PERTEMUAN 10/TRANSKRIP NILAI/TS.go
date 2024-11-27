package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
    if n <= 4 && n >= 3.00  {
            var tak float64
            fmt.Scan(&tak)
         if tak < 2.00 {
            fmt.Println("Mahasiswa Lulus Cumlaude dengan IPK",n,"Dan Status predikat TAK Poor")
        }else if tak >= 2.00 && tak <= 2.75 {
            fmt.Println("Mahasiswa Lulus Cumlaude dengan IPK",n,"Dan Status predikat TAK Fair")
        }else if tak >= 2.76 && tak <= 3.00 {
            fmt.Println("Mahasiswa Lulus Cumlaude dengan IPK",n,"Dan Status predikat TAK Satisfactory")
        }else if tak >= 3.01 && tak <= 3.50 {
            fmt.Println("Mahasiswa Lulus Cumlaude dengan IPK",n,"Dan Status predikat TAK Very good")
        }else {
            fmt.Println("Mahasiswa Lulus Cumlaude dengan IPK",n,"Dan Status predikat TAK Excellent")
        }
        
    } else {
        fmt.Println("Mahasiswa tidak cumlaude dan tidak bisa menginputkan tak")
    } 
}

// INI WM FAZA 