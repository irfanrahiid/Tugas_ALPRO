package main

import "fmt"

func main() {
	var sisi, volume float64

	fmt.Print("masukan panjang sisi = ")
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Println("volume kubus adalah = ", volume)
}