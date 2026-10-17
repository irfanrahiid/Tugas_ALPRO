//program ini saya buat untuk menghitung luas dan keliling persegi panjang

package main

import "fmt"

func main() {
    var panjang,lebar,luas,keliling float64

	fmt.Print("masukan panjang =")
    fmt.Scan(&panjang)
	fmt.Print("masukan lebar =")
	fmt.Scan(&lebar)


	luas = panjang * lebar 
	keliling = 2 * panjang + 2 * lebar
	
	fmt.Println("luas persegi panjang =", luas)
	fmt.Println("leliling persegi panjang =", keliling)

}