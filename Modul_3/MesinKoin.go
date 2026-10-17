package main

import "fmt"

func main() {
	var koin, Kemas, Kperak, Ktembaga int

	fmt.Print("masukan jumlah koin yang mau di konversi = ")
	fmt.Scan(&koin)

	Kemas = koin / 9
	Kperak = koin % 9 / 3
	Ktembaga = koin % 9 % 3

	fmt.Printf("koin yang kamu miliki setelah di konversi adalah %d koin emas, %d koin perak, dan %d koin tembaga", Kemas, Kperak, Ktembaga)
} 