package main

import "fmt"

func main() {
	var n, sisa, jam, menit, detik int

	fmt.Print("masukan detik = ")
	fmt.Scan(&n)

	jam = n / 3600
	sisa = n % 3600
	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf(" %d jam, %d menit, %d detik", jam, menit, detik)
}