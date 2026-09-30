/*Program ini saya buat untuk mengkonversi suhu dari celcius ke reamur, fahrenheit,
dan kelvin, pengguna cukup memasukan nilai celciusnya saja dan program akan otomatis menghitung
semuanya dan mengeluarkan hasilnya di output*/

package main

import "fmt"

func main() {
	var c int

	fmt.Print("masukan nilai celcius = ")
	fmt.Scan(&c)

	r := c * 4 / 5
	f := (c * 9 / 5) + 32
	k := c + 273

	fmt.Println("nilai reamur = ", r)
	fmt.Println("nilai fahrenheit = ", f)
	fmt.Println("nilai kelvin = ", k)
}