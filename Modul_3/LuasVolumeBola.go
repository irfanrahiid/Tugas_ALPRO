/*Program ini saya buat untuk menghitung luas dan volume bola, pengguna
cukup memasukan jari-jari bola dan program akan otomatis menghitung dan mengeluarkan outputnya*/

package main

import "fmt"

func main() {
	var r int

	fmt.Print("masukan nilai jari-jari lingkaran = ")
	fmt.Scan(&r)
	const pi = 3.1415926535

	var vbola, lbola float64

	vbola = (4.0 / 3.0) * pi * float64(r * r * r)
	lbola = 4 * pi * float64(r * r)

	fmt.Print("volume bola = ", vbola)
	fmt.Println("luas bola = ", lbola)
}