/*
	program TampilBilangan
	kamus
		n1, n2, i int
	algoritma
		input(n1, n2)
		for i <- n1 to n2 do
			output(i)
		endfor
	endprogram
*/

package main

import "fmt"

func main() {
	var n1, n2, i int
	
	fmt.Print("masukan angka pertama = ")
	fmt.Scan(&n1)
	fmt.Print("masukan angka kedua = ")
	fmt.Scan(&n2)

	for i = n1; i <= n2; i += 2 {
		fmt.Print(i)
	} 
}