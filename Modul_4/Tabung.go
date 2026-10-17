package main

import "fmt"

const PI float64 = 3.14
type tabung struct{
	tinggi, jari int
	luas, volume float64
	luasalas, luasdinding float64
}

func main() {
	var t tabung

	fmt.Print("masukan jari-jari = ")
	fmt.Scan(&t.jari)

	fmt.Print("masukan tinggi = ")
	fmt.Scan(&t.tinggi)

	t.luasalas = PI * (float64(t.jari) * float64(t.jari))
	t.luasdinding = float64(t.tinggi) * (2 * PI * float64(t.jari))
	t.luas = 2 * t.luasalas + t.luasdinding
	t.volume = t.luasalas * float64(t.tinggi)

	fmt.Printf("luas = %.2f \nvolume = %.2f", t.luas, t.volume)
}