package main

import "fmt"

func main() {
	var n, sisa, jam, menit, detik int

	fmt.Print("masukan nilai detik = ")
	fmt.Scan(&n)

	jam = n / 3600
	sisa = n % 3600
	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf("waktu rapat anda %d jam, %d menit, %d detik", jam, menit, detik)
}