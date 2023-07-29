package main

import "fmt"

func slice() {
	var slice = []string{"mangga", "manggis", "madu", "madura", "makasar"}
	newSlice := slice[0:5]
	newSlice1 := slice[0:0]
	newSlice2 := slice[0:]
	newSlice3 := slice[:5]
	newSlice4 := slice[:]

	fmt.Println(newSlice)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)
	fmt.Println(newSlice4)
}
func sliceReference() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	aFruits := fruits[0:3]
	bFruits := fruits[1:2]

	aaFruits := aFruits[1:2]
	baFruits := bFruits[0:1]

	fmt.Println("===Fruits awal===")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
	fmt.Println("===setelah grape diganti pinapple===")

	baFruits[0] = "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	var (
		nama string
		umur int
	)
	nama = "sitamvan"
	umur = 1
	fmt.Println(nama, umur)
}

func main() {
	fmt.Println("=====Ini Slice normal=====")
	slice()
	fmt.Println("=====Contoh Slice Reference=====")
	sliceReference()
}
