//program ini saya buat untuk menghitung karat pada emas

package main

import "fmt"

func main() {
    var gram,karat float64
    fmt.Println("input gram =")
    fmt.Scan(&gram)
    karat = gram + 4 * gram
    fmt.Println("karat =", karat)
}