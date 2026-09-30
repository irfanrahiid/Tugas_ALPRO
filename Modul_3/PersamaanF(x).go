/*Program ini saya buat untuk menghitung persamaan nilai f(x), pengguna
cukup memasukan nilai x dan program akan otomatis menghitung dan mengeluarkan hasilnya di output*/

package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan nilai x = ")
	fmt.Scan(&x)

	var f float64

	f = 2 / (float64(x) + 5) + 5
	fmt.Println("nilai f(x) = ", f)
}