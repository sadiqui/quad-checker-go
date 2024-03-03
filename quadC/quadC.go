package main

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				fmt.Print("A")
			} else if row == 1 && col > 1 && col < x {
				fmt.Print("B")
			} else if row == 1 && col == x {
				fmt.Print("A")
			} else if row > 1 && row < y && col == 1 {
				fmt.Print("B")
			} else if row == y && col == 1 {
				fmt.Print("C")
			} else if row == y && col > 1 && col < x {
				fmt.Print("B")
			} else if row == y && col == x {
				fmt.Print("C")
			} else if row > 1 && row < y && col == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	QuadC(5, 3)
	QuadC(5, 1)
	QuadC(1, 1)
	QuadC(1, 5)
}

// ABBBA
// B   B
// CBBBC
// ABBBA
// A
// A
// B
// B
// B
// C
