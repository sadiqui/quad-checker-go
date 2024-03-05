package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	if len(os.Args) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])
		QuadC(x, y)
	} else {
		return
	}
}
