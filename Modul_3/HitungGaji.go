package main

import "fmt"

func main() {
	var gp, jl, bl, pt, gb int

	fmt.Print("masukan gaji pokok anda = ")
	fmt.Scan(&gp)
	fmt.Print("masukan jumlah jam lembur anda = ")
	fmt.Scan(&jl)

	bl = jl * 45000
	pt = (55 * gp) / 1000
	gb = gp + bl - pt

	fmt.Println("gaji bersih anda sebesar =", gb)
}