/*
program BMI
kamus
   berat, tinggi, bmi : real
algoritma
   input (berat, tinggi)
   bmi = berat div (tinggi * tinggi)
   output (bmi)
endprogram
*/

package main

import "fmt"

func main () {
	var berat, tinggi, bmi float64

	fmt.Print("masukan berat badan anda = ")
	fmt.Scan(&berat)

	fmt.Print("masukan tinggi badan anda = ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Printf("BMI anda = %.2f", bmi)
}