package main

import "fmt"

func main() {
	var dalam, tekanan, batas float64
	
	fmt.Print("masukan nilai kedalaman = ")
	fmt.Scan(&dalam)

	fmt.Print("masukan batas tekanan maksimal = ")
	fmt.Scan(&tekanan)

	batas = float64(1025) * 9.8 * dalam

	var aman bool
	aman = tekanan <= batas
	
	fmt.Println("besaran tekanan = ", batas)
	fmt.Print("apakah aman untuk penyelam = ", aman)
}