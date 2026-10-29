package main

import "fmt"

func main() {
	var x, y, z, jum int

	fmt.Print("masukan bilangan kecil = ")
	fmt.Scan(&x)
	fmt.Print("masukan bilangan besar = ")
	fmt.Scan(&y)

	jum = 0

	for z = x; z <= y; z++ {
		jum = jum + z
	}
	fmt.Print("hasil penjumlahan semua bilangan kecil ke besar adalah = ", jum)
}