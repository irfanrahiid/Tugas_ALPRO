package main

import "fmt"

func main() {
	var aeos int
	var hasil bool

	fmt.Print("masukan jumlah koin aeos = ")
	fmt.Scan(&aeos)
	
	hasil = aeos >= 1500 
	fmt.Print("apakah jumlah koin aeos anda cukup = ", hasil)
}