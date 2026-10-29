package main

import "fmt"

func main () {
	var n,h, i, b1, b2, b3, b4, b5, b6 int

	fmt.Print("masukan banyak baris = ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++{
		fmt.Println("masukan baris bilangan (gunakan spasi perhuruf)")
		fmt.Scan(&b1, &b2, &b3, &b4, &b5, &b6)
		h = b1 + b2 - b3 * b4 / b5 % b6
		fmt.Println("hasil perhitungan = ", h)
	}
}