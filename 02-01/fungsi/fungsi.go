package main

import "fmt"

func main() {

	var coba2 [2]int = [2]int{1, 2}
	fmt.Println(coba2[0] + coba2[1])

	fmt.Println(coba([]string{"menteng"}, []int{1}))

}
func coba(nama []string, umur []int) ([]string, []int) {

	return nama, umur
}
