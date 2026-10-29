package main

import "fmt"

func main() {
	var n, i int
	var jam, menit, sisa, durasi int

	fmt.Print("masukan jumlah film yang di tonton = ")
	fmt.Scan(&n)

	durasi = 0

	for i = 1; i <= n; i++ {
		fmt.Print("masukan menit = ")
		fmt.Scan(&menit)

		durasi = durasi + menit
	}

	jam = durasi / 60
	sisa = durasi % 60
	menit = sisa

	fmt.Printf("total durasi menonton film anda = %d jam, %d menit", jam, menit)
}