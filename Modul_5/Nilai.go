package main
import "fmt"

func main (){
	const (
		kuis float64 = 0.15
		tugas float64 = 0.15
		praktikum float64 = 0.2
		UTS float64 = 0.25
		UAS float64 = 0.25
	)

	type mahasiswa struct {
		NIM string
		Nama string
		NilaiTugas float64
		NIilaiKuis float64
		NilaiPraktikum float64
		NilaiUAS float64
		NilaiUTS float64
		NilaiAkhir float64
	}

	var DataMahasiswa mahasiswa
	var nilaiAkhir float64

	fmt.Scan(&DataMahasiswa.Nama, &DataMahasiswa.NIM, &DataMahasiswa.NilaiTugas, &DataMahasiswa.NIilaiKuis, &DataMahasiswa.NilaiPraktikum, &DataMahasiswa.NilaiUAS, &DataMahasiswa.NilaiUTS)

	nilaiAkhir = DataMahasiswa.NilaiTugas * tugas + DataMahasiswa.NIilaiKuis * kuis + DataMahasiswa.NilaiPraktikum * praktikum + DataMahasiswa.NilaiUTS * UTS + DataMahasiswa.NilaiUAS * UAS

	fmt.Println("Data Mahasiswa")
	fmt.Println("Nama: = ", DataMahasiswa.Nama)
	fmt.Println("NIM: ", DataMahasiswa.NIM)
	fmt.Println("Nilai Tugas: ", DataMahasiswa.NilaiTugas)
	fmt.Println("Nilai Kuis: ", DataMahasiswa.NIilaiKuis)
	fmt.Println("Nilai Praktikum: ", DataMahasiswa.NilaiPraktikum)
	fmt.Println("Nilai UAS: ", DataMahasiswa.NilaiUAS)
	fmt.Println("Nilai UTS: ", DataMahasiswa.NilaiUTS)
	fmt.Println("Nilai Akhir: ", nilaiAkhir)


}