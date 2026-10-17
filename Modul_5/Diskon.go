package main

import "fmt"

func main () {
	 var hargaawal float64

	 fmt.Scan(&hargaawal)

	 diskon1 := hargaawal * 0.90
	 hargaakhir := diskon1 * 0.80

	 fmt.Printf("%.2f", hargaakhir)
}