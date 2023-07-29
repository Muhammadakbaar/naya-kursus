package main

import (
	"fmt"
)

func cobaSlice() {
	var anu1 = []string{"satu", "dua", "tiga", "empat"}
	anubaru := anu1[0:4]
	fmt.Println(anubaru)
	anu1[0] = "lima"
	fmt.Println(anu1[0])
	fmt.Println(cap(anu1))
	anubarulagi := anu1[0:3]
	fmt.Println(cap(anubarulagi))

}
func mencoba() {
	var fruits = []string{"satu", "dua", "tiga", "empat", "lima"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
func pakeMake() {
	var fruits = make([]int, 3)
	fruits[0] = 20
	fruits[1] = 30
	fruits[2] = 30
	for i := 0; i < 5; i++ {
		fmt.Println(fruits)
	}
}

func main() {
	//cobaSlice()
	//mencoba()
	pakeMake()
}
