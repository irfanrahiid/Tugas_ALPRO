/*Program ini saya buat untuk mengkonversi suhu dari Fahrenheit ke celcius,
kita bisa memasukan nilai suhu fahrenheit pada input dan program akan otomatis
mengkoversi nilai suhu fahrenheit dan mengeluarkan nilai suhu celcius pada output*/

package main

import "fmt"

func main() {
	var cel, fah int64

	fmt.Print("masukan fahrenheit = ")
	fmt.Scanln(&fah)

	cel = (fah - 32) * 5 / 9

	fmt.Printf("hasil = %d celcius", cel)
}