/*Program ini saya buat untuk menghitung luas dan keliling Lingkaran,
kita bisa memasukan nilai jari-jari lingkaran pada input dan program akan
otomatis menghitung dan mengeluarkan nilainya pada output*/

package main

import "fmt"

func main() {
	var r,luas,keliling float64
	phi := 3.14

	fmt.Print("masukan jari-jari = ")
	fmt.Scanln(&r)

	luas = phi * r * r
	keliling = 2 * phi * r

	fmt.Println("luas lingkaran = ", luas)
	fmt.Print("Keliling lingkaran = ", keliling)
}