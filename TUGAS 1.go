package main

import "fmt"

type mahasiswa struct {
	nama   string
	alasan string
}

func main() {
	var data1 = mahasiswa{}
	data1.nama = "Denada Ramschie
	data1.alasan = "Menambah wawasan ilmu yang baru"

	var data2 = mahasiswa{}
	data2.nama = "Luqman Fauzi"
	data2.alasan = "Ingin mempunyai skill khusus pada bagian back end"

	var data3 = mahasiswa{}
	data3.nama = "Haiqal Malik"
	data3.alasan = "Mengembangkan skill"

	var data4 = mahasiswa{}
	data4.nama = "Ruben"
	data4.alasan = "Menambah pengetahuan"

	fmt.Println("Nama :", data1.nama)
	fmt.Println("Alasan :", data1.alasan)

	fmt.Println("Nama :", data2.nama)
	fmt.Println("Alasan :", data2.nama)

	fmt.Println("Nama :", data3.nama)
	fmt.Println("Alasan :,data3.nama")

}
