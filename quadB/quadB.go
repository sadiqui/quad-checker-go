package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				fmt.Print("/")
			} else if row == 1 && col > 1 && col < x {
				fmt.Print("*")
			} else if row == 1 && col == x {
				fmt.Print("\\")
			} else if col == 1 && row > 1 && row < y {
				fmt.Print("*")
			} else if row == y && col == 1 {
				fmt.Print("\\")
			} else if row == y && col > 1 && col < x {
				fmt.Print("*")
			} else if row == y && col == x {
				fmt.Print("/")
			} else if col == x && row > 1 && row < y {
				fmt.Print("*")
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
		QuadB(x, y)
	} else {
		return
	}
}
