package main

import (
	"fmt"
)

func main() {
	var angka1 int
	var angka2 int

	for {
		fmt.Print("Masukan angka pertama = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan angka kedua = ")
		fmt.Scan(&angka2)
		if angka1 == angka2 {
			fmt.Println("benar")
		} else if angka1 != angka2 {
			fmt.Println("salah")
		}

		fmt.Print("Masukan nilai pertama = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan nilai kedua = ")
		fmt.Scan(&angka2)
		if angka1 <= angka2 {
			fmt.Println("angka pertama lebiih kecil dari angka kedua")
		} else if angka1 >= angka2 {
			fmt.Println("angka pertama lebih besar dari angka kedua")
		}
		break
	}

}
