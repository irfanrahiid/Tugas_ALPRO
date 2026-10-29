package main

import "fmt"

func main() {
	var n, i, saldo, kenaikan, total, penambahan int
	
	fmt.Print("masukan jumlah minggu = ")
	fmt.Scan(&n)
	fmt.Print("masukan saldo awal = ")
	fmt.Scan(&saldo)
	fmt.Print("masukan jumlah kenaikan setiap minggu = ")
	fmt.Scan(&kenaikan)

	total = saldo
	penambahan = kenaikan

	for i = 1; i <= n; i++ {
		total = total + penambahan 
		penambahan = penambahan + kenaikan
	}

	fmt.Print("total saldo anda = ", total)
}