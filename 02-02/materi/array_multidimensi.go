package main

import (
	"fmt"
)

func perulanganArray2(coba []int, coba2 []string) ([]int, []string) {

	return coba, coba2
}
func perulanganArrayForRange() {
	var fruits = [4]string{"mangga", "jambu", "manggis", "rambutan"}
	for i, fruit := range fruits {
		fmt.Println("isinya perulangan range", i, fruit)
	}
}
