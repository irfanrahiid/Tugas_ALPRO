package main

import "fmt"

func main() {
	var i, n, ganjil int

	fmt.Print("masukan jumlah bilangan bulat yang ingin di tampilkan = ")
	fmt.Scan(&n)


	for i = 1; i <= n; i++ {
		ganjil = 2*i - 1
		fmt.Println(ganjil)
	}
}