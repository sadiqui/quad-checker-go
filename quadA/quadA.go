package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				fmt.Print("o")
			} else if row == 1 && col > 1 && col < x {
				fmt.Print("-")
			} else if row == 1 && col == x {
				fmt.Print("o")
			} else if row > 1 && row < y && col == 1 {
				fmt.Print("|")
			} else if row == y && col == 1 {
				fmt.Print("o")
			} else if row == y && col > 1 && col < x {
				fmt.Print("-")
			} else if row == y && col == x {
				fmt.Print("o")
			} else if row > 1 && row < y && col == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}

// o---o
// |   |
// o---o
// o---o
// o
// o
// |
// |
// |
// o
