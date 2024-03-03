package main

import "fmt"

func QuadD(x, y int) {
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
			} else if row > 1 && row < y && col == 1 {
				fmt.Print("B")
			} else if row == y && col == 1 {
				fmt.Print("A")
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
	QuadD(5, 3)
	QuadD(5, 1)
	QuadD(1, 1)
	QuadD(1, 5)
}

// ABBBC
// B   B
// ABBBC
// ABBBC
// A
// A
// B
// B
// B
// A
