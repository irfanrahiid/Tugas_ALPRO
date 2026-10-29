/*
	program hitungpangkat
	kamus
		n1, n2, i, j : int
	algoritma
		input(n1, n2)
		j = 1
		for i <- i to n2 do
			j = j * n1
		endfor
		output(j)
	endprogram
*/

package main

import "fmt"

func main() {
	var n1, n2, i, j int

	fmt.Print("masukan nilai yang ingin di pangkatkan = ")
	fmt.Scan(&n1)
	fmt.Print("masukan nilai pangkatnya = ")
	fmt.Scan(&n2)

	j = 1

	for i = 1; i <= n2; i++ {
		j = j * n1
	}

	fmt.Print(j)
}