package main

import "fmt"

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
func main() {
	cobaSlice()
}
