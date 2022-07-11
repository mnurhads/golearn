package main

import (
	"fmt"
)

// function 
func LuasSegitiga(data1, data2 int) int {
	return (data1 * data2) / 2
}

func PanjangList(array []int) {
	fmt.Println("\n", len(array), "List Data")
}

func main() {
	var angka int 
	var angka2 int
	var hasil int
	var comment string
	// auto :=
	ckd := "Hello,"
	// string call
	comment = "Ini hasil dari penjumlahan yang tersedia"

	// constan otomatis akan mengunci file
	const firstname = "Budi Budarsono"

	angka = 10
	angka2 = 30

	hasil = angka * angka2
	// tampilkan
	fmt.Println(ckd, firstname, ". Ini Tes Connection, World => angka :", hasil, comment)
	fmt.Println("===================================================================================")
	// Aritmatika => Kalkulator sederhana
	var namaBarang string = "Macbook Pro M1 Chipset"
	var uangMuka int = 1000000 // uang muka yang diberikan pembeli
	var cicilan int  = 24 // jumlah cicilan yang disetujui
	var cicilLancar int = 5 // jumlah bulan cicilan yang berjalan
	var price int  = 25000000 // harga normal barang
	var bunga int   = 3 // dalam konversi decimal 5%
	var bungaDown int = 2 // dalam 13 bulan masa cicilan, akan diberlakuakn penurunan jumlah bunga
	// hitung pengurangan DP
	var uangSisa int = price - uangMuka
	// hitung bunga perbulan yang dibebankan
	var hargaTotal int
	var bungaMonth int
	fmt.Println(" -- Perhitungan Jumlah Cicilan Sederhana--")
	fmt.Println("Nama Barang  :", namaBarang)
	fmt.Println("Harga Barang : Rp.", price, "| Cicilan  :", cicilan, "/Bulan")
	fmt.Println("Uang Muka    : Rp.", uangMuka, " | Hutang   : Rp.", uangSisa)
	// ubah intger to decimal
	bungaMonth = bunga * uangSisa / 100
	fmt.Println("Bunga(",bunga,"%  - ",bungaDown,"%)   : Rp.", bungaMonth, "Potongan Rp", (bungaDown * bungaMonth / 100) )
	hargaTotal = uangSisa / cicilan + bungaMonth
	fmt.Println("Jumlah Cicilan :", cicilLancar, "Kali| Sisa Cicilan : ", cicilan - cicilLancar,"Kali")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Total yang harus dicicil setiap bulan adalah : Rp.", hargaTotal)
	fmt.Println("==============================================================================")
	// perulangan golang dasar
	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println("*", " ")
		}

		fmt.Println()
	}

	// matriks perulangan golang
	var matrik [][]int32 = [][]int32{
		{1,2},
		{3,4},
		{5,6},
		{7,8},
	}

	fmt.Println("Speak Dadu Up")
	for i := 0; i < len(matrik); i++ {
		for new := 0; new < len(matrik[i]); new++ {
			fmt.Println("Random Dadu", i, new)
			fmt.Println("List Random", matrik[i][new])
		}
	}
	fmt.Println("=================================================")
	// call a function simple
	a := LuasSegitiga(20, 10)
	fmt.Println("Luas Segitika adalah,rumus (Alas x Tinggi) / 2, yaitu:", a)

	arr := []int{1, 3, 5, 10, -5, -7, 22, 18, 20}
	PanjangList(arr)
	
}