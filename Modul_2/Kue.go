//program ini saya buat untuk menghitung sisa kue untuk si bungsu

package main

import "fmt"

func main() {
	var y, x int
	fmt.Print("Masukkan jumlah kue (y): ")
	fmt.Scanln(&y)
	fmt.Print("Masukkan jumlah anggota keluarga (x): ")
	fmt.Scanln(&x)
	sisaKue := y % x
	fmt.Printf("Jumlah sisa kue untuk anak bungsu: %d\n", sisaKue)
}