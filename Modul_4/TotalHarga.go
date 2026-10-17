package main

import "fmt"

func main() {
	var harga, diskon, total int

	fmt.Print("masukan total harga = ")
	fmt.Scan(&harga)

	fmt.Print("masukan diskon = ")
	fmt.Scan(&diskon)

	total = harga - (harga * diskon / 100)

	fmt.Print("total belanja akhir = ", total)
}