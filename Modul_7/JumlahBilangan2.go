package main

import "fmt"

func main() {
	var x, y, n int
	var jumlah float64

	fmt.Print("masukan nilai bilangan kecil = ")
	fmt.Scan(&x)

	fmt.Print("masukan nilai bilangan besar = ")
	fmt.Scan(&y)

	n = (y - x) + 1
	jumlah = float64(n * (x + y)) / 2.0

	fmt.Print("hasil penjumlahan semua bilangan dari kecil ke besar adalah = ", jumlah)
}