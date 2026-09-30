/*Program ini saya buat untuk menerima tiga input string dari pengguna,
menampilkannya, lalu menggeser nilai variable tersebut secara melingkar ke kiri*/

package main

import "fmt"

func main() {    
	var (
		satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
} 