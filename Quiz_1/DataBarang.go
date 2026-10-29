package main

import "fmt"

type barang struct {
	nama string
	harga int64
	stok int64
}


func main() {
	var Jkategori int64
	var Jbarang int64
	var i,k int64
	var b barang

	fmt.Print("masukan jumlah kategori barang = ")
	fmt.Scanln(&Jkategori)
	fmt.Println(" ")

	for i = 1; i <= Jkategori; i++ {
		fmt.Print("masukan jumlah barang yang ingin di input = ")
		fmt.Scanln(&Jbarang)
		fmt.Println(" ")

		for k = 1; k <= Jbarang; k++ {
			fmt.Print("masukan nama barang = ")
			fmt.Scanln(&b.nama)

			fmt.Print("masukan harga barang = ")
			fmt.Scanln(&b.harga)

			fmt.Print("masukan jumlah stok barang = ")
			fmt.Scanln(&b.stok)
			fmt.Println(" ")

			fmt.Println("BERHASIL INPUT :")

			fmt.Print("nama barang = ")
			fmt.Println(b.nama)

			fmt.Print("harga barang = Rp.")
			fmt.Println(b.harga)

			fmt.Print("jumlah stok = ")
			fmt.Println(b.stok)
			fmt.Println(" ")
		}
	}
}

