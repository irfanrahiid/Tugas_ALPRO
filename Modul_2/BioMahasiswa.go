/*Program ini saya buat untuk menginput data mahasiswa, kita bisa memasukan
nilai nama, nim, dan kelas pada input dan program akan otomatis Menyusun data
dan mengeluarkan data yang di input pada output*/

package main

import "fmt"

func main() {
	var nama, kelas string
	var nim int64

	fmt.Print("input nama = ")
	fmt.Scanln(&nama)

	fmt.Print("input nim = ")
	fmt.Scanln(&nim)

	fmt.Print("input kelas = ")
	fmt.Scanln(&kelas)

	fmt.Printf("perkenlkan nama saya %s, salah satu mahasiswa prodi S1-IF dari kelas %s dengan NIM %d.", nama, kelas, nim)

}