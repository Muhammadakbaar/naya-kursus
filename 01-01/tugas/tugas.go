package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		var number int
		var err error

		fmt.Print("Masukan angka = ")
		fmt.Scanln(&input)

		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input harus berupa angka!")
			continue
		}
		factorial := 1
		for i := 1; i <= number; i++ {
			factorial *= i
		}
		fmt.Printf("Faktorial dari %d adalah = %d \n", number, factorial)
		break
	}
}
