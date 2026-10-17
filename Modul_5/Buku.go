package main

import "fmt"

type Buku struct {
	judul       string
	genre       string
	tahunRilis  string
	penilaian   float64
}

func main() {

	var buku1, buku2 Buku

	fmt.Scan(&buku1.judul, &buku2.judul)
	fmt.Scan(&buku1.genre, &buku2.genre)
	fmt.Scan(&buku1.tahunRilis, &buku2.tahunRilis)
	fmt.Scan(&buku1.penilaian, &buku2.penilaian)

	fmt.Println("--- Book 1 ---")
	fmt.Println(buku1.judul)
	fmt.Println(buku1.genre)
	fmt.Println(buku1.tahunRilis)
	fmt.Println(buku1.penilaian)

	fmt.Println("--- Book 2 ---")
	fmt.Println(buku2.judul)
	fmt.Println(buku2.genre)
	fmt.Println(buku2.tahunRilis)
	fmt.Println(buku2.penilaian)
}