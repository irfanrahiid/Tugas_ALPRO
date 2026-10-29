package main

import "fmt"

func main() {
	var dinar, dirham, fals, qirat int

	fmt.Print("masukan jumlah uang qirat = ")
	fmt.Scan(&qirat)

	dinar =  qirat / 600
	dirham = (qirat % 600) / 60
	fals = ((qirat % 600) % 60) / 6
	qirat = ((qirat % 600) % 60) % 6

	fmt.Println("jumlah uang dinar anda = ", dinar)
	fmt.Println("jumlah uang dirham anda = ", dirham)
	fmt.Println("jumlah uang fals anda = ", fals)
	fmt.Println("sisa uang qirat anda = ", qirat)
}