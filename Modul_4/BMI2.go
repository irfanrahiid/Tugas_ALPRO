package main

import "fmt"

func main () {
	var tinggi, bmi, berat float64

	fmt.Print("masukan BMI anda = ")
	fmt.Scan(&bmi)

	fmt.Print("masukan tinggi badan anda = ")
	fmt.Scan(&tinggi)

	berat = bmi * tinggi * tinggi

	fmt.Printf("berat badan anda anda = %.0f\n", berat)
}