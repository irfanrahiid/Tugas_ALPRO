package main

import "fmt"

func main() {
	var rupiah int
	fmt.Scan(&rupiah)

	var dolar float64
	dolar = float64(rupiah) / 15000
	fmt.Printf("USD %.2f", dolar) 
}