package main

import "fmt"

func main() {
	var n, j, x, total int

	fmt.Println("masukan jumlah perenang = ")
	fmt.Scan(&n)
	
	for j = 1; j <= n; j++ {

		fmt.Print("masukan jarak yang di tempuh oleh para perenang = ")
		fmt.Scan(&x)

		total = x * 5
		
		fmt.Println("total jarak yang di tempuh oleh setiap perenang dalam 5 putaran = ", total)
	}
}