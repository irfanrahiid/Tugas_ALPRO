//program ini saya buat untuk menghitung penjumlahan bilangan bulat

package main

import "fmt"

func main() {
    var a, b, c, d, jumlah int
	a = 2
	b = 4
	c = 6
	d = 8
    fmt.Scan(&a, &b, &c, &d)
    jumlah = a + b + c + d
    fmt.Println(jumlah)
}