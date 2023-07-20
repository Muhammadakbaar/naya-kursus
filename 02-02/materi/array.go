package main

import "fmt"

func getArray() {
	var nama = [3]string{
		"Muhammad",
		"Akbar",
		"Hakim",
	}
	var umur [2]int
	umur[0] = 21
	umur[1] = 26

	fmt.Println(nama)
	fmt.Println("Mengacak index array", nama[1], nama[2], nama[0])
	fmt.Println(umur)

}
func perulanganArray() {

	var mengulangArray = [3]string{"satu", "dua", "tiga"}
	for i, ulang := range mengulangArray {
		fmt.Println("isinya", i, ulang)

	}
	for i := 0; i < len(mengulangArray); i++ {
		fmt.Println(i, mengulangArray)
	}

}
func arrayTitikTiga() {
	var anu = [3][4]int{{2, 3, 4}, {3, 5, 6, 7}, {1}}
	fmt.Println(anu)
	fmt.Println(anu[1][2])

}
func main() {
	getArray()
	perulanganArray()
	arrayTitikTiga()
}
