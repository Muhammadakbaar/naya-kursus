package main

import "fmt"

func main() {
	var angka1 int = 10
	var angka2 int = 5

	fmt.Println("==Operator Aritmatika==")
	fmt.Println(angka1, "+", angka2, "=", angka1+angka2)
	fmt.Println(angka1, "-", angka2, "=", angka1-angka2)
	fmt.Println(angka1, "x", angka2, "=", angka1*angka2)
	fmt.Println(angka1, "/", angka2, "=", angka1/angka2)
	fmt.Println(angka1, "%", angka2, "=", angka1%angka2)

	fmt.Println("==Operator Assignment==")
	angka1 += 5
	fmt.Println(angka1)
	angka1 -= angka2
	fmt.Println(angka1)
	angka1 *= angka2
	fmt.Println(angka1)
	angka1 /= angka2
	fmt.Println(angka1)
	angka1 %= angka2
	fmt.Println(angka1)

	fmt.Println("==Operator Perbandingan==")
	fmt.Println("angka1 sama dengan angka2 =", angka1 == angka2)
	fmt.Println("angka1 tidak dengan angka2 =", angka1 != angka2)
	fmt.Println("angka1 lebih kecil dari angka2 =", angka1 < angka2)
	fmt.Println("angka1 lebih kecil atau sama dengan angka2 =", angka1 <= angka2)
	fmt.Println("angka1 lebih besar dari angka2 =", angka1 > angka2)
	fmt.Println("angka1 lebih besar atau sama dengan angka2 =", angka1 >= angka2)

	fmt.Println("==Operator Logika==")
	angka3 := true
	angka4 := false
	fmt.Println("angka3 atau angka4 = ", angka3 || angka4)
	fmt.Println("angka3 dan angka4 = ", angka3 && angka4)
	fmt.Println("angka3 tidak angka4 = ", angka3, !angka4)

}
