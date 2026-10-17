package main

import "fmt"

func main() {
	const harga int = 20000
	const pajak float64 = 0.11

	var jumlah int

	fmt.Print("masukan jumlah kopi = ")
	fmt.Scan(&jumlah)

	totalharga := float64(jumlah * harga)
	totalpajak := totalharga * pajak
	totalbayar := int(totalharga + totalpajak)

	fmt.Print("total harga yang harus anda bayar adalah = ", totalbayar)
}