/*
	program hitungfaktorial
	kamus
		n1, i, j : int
	algoritma
		input(n1)
		j = 1
		for i <- 1 to n1 do
			j = j * i
		endfor
		output(j)
	endprogram
*/

package main

import "fmt"

func main() {
	var n1, i, j int
	
	fmt.Print("masukan nilai = ")
	fmt.Scan(&n1)

	j = 1
	
	for i = 1; i <= n1; i++ {
		j = j * i
	}

	fmt.Print(j)
}