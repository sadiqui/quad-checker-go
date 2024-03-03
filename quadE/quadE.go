package main

import "fmt"

func QuadE(x, y int) {
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
				fmt.Print("C")
			} else if col == 1 && row > 1 && row < y {
				fmt.Print("B")
			} else if row == y && col == 1 {
				fmt.Print("C")
			} else if row == y && col > 1 && col < x {
				fmt.Print("B")
			} else if row == y && col == x {
				fmt.Print("A")
			} else if col == x && row > 1 && row < y {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	QuadE(5, 3)
	QuadE(5, 1)
	QuadE(1, 1)
	QuadE(1, 5)
}

// ABBBC
// B   B
// CBBBA
// ABBBC
// A
// A
// B
// B
// B
// C
