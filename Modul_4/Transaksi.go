package main

import "fmt"

type data struct {
	nama string
	jumlah int
	harga, total float64
}

func main() {
	var d data

	fmt.Print("masukan nama barang = ")
	fmt.Scan(&d.nama)

	fmt.Print("masukan jumlah barang = ")
	fmt.Scan(&d.jumlah)

	fmt.Printf("masukan harga satuan = Rp.")
	fmt.Scan(&d.harga)

	d.total = float64(d.jumlah) * d.harga

	fmt.Printf("total harga barang anda = Rp. %.2f", d.total)
}