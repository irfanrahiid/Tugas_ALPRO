/*Program ini saya buat untuk mengecek apakah tahun yang di input oleh pengguna termasuk
tahun kabisat atu bukan, jika iya hasilnya akan true dan jika tidak hasilnya akan fals*/

package main

import "fmt"

func main() {
	var tahun int64

	fmt.Print("masukan tahun = ")
	fmt.Scan(&tahun)

	kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Println("tahun kabisat = ", kabisat)
}