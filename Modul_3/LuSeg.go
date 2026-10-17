package main

import "fmt"

func main () {
	var pal, tiseg int

	fmt.Print("masukan panjang alas = ")
	fmt.Scan(&pal)

	fmt.Print("masukan tinggi segitiga = ")
	fmt.Scan(&tiseg)

	palflt := float64(pal)
	tisegflt := float64(tiseg)

	luas := palflt * tisegflt / 2

	fmt.Println("luas segitiga adalah =", luas)
}