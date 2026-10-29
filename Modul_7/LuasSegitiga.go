package main

import "fmt"

func main() {
	var j, alas, tinggi, n int
	var luas float64

	fmt.Print("masukan banyak segitiga = ")
	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		fmt.Print("masukan alas = ")
		fmt.Scan(&alas)

		fmt.Print("masukan tinggi = ")
		fmt.Scan(&tinggi)

		luas = 0.5 * float64(alas * tinggi)

		fmt.Println("luas segitiga = ", luas)
	}
}