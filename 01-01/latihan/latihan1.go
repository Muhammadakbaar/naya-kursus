package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 5
	fmt.Println("==sebelum==")
	fmt.Println(" a = ", a)
	fmt.Println(" b = ", b)
	fmt.Println("==sesudah==")
	a, b = b, a
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
