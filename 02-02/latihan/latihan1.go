package main

import "fmt"

func main() {
	// Membuat variabel nilai dengan tipe data array
	nilai := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var index int
	fmt.Print("Masukan index dari 0 sampai 9 = ")
	fmt.Scan(&index)
	nilaiIndex := nilai[index]

	if nilaiIndex < 5 {
		fmt.Printf("Nilai index %d: Tidak Lulus\n", index)
	} else {
		fmt.Printf("Nilai index %d: Lulus\n", index)
	}

}
