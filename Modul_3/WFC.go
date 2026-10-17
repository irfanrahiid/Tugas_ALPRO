package main

import "fmt"

func main() {
	var N, hari int
	
	fmt.Print("masukan total jumlah hari = ")
	fmt.Scan(&N)

	namaHari := [8]string{"", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	hari = (4+N-1)%7 + 1

	fmt.Printf("%d hari setelah hari Kamis (4) adalah %s (%d)\n", N, namaHari[hari], hari)
}