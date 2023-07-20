package main

func main() {
	for i := 5; i-1 > 0; i-- {
		for j := i; j <= 5; j++ {
			print(" ")
		}

		for j := i; j > 0; j-- {
			print("* ")
		}
		println()
	}

	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			print(" ")
		}

		for j := 0; j <= i; j++ {
			print("* ")
		}
		println()
	}

}
