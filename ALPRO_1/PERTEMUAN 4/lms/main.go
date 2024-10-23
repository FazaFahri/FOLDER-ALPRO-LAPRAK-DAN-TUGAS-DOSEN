package main
import"fmt"
func main() {
    var huruf int
    var status bool
    fmt.Scanf("%c",&huruf)
    status = 'a' == huruf || 'i' == huruf || 'u' == huruf || 'e' == huruf || 'o' == huruf
  fmt.Print(status)
}