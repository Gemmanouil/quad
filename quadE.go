package main

import "fmt"

func QuadE(x, y int) { // checks if x,y positive (if not exit)
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ { // loop for rows
		for col := 0; col < x; col++ { // loop for columns
			if row == 0 && col == 0 { // checks point (0,0)
				fmt.Print("A")
			} else if row == 0 && col == x-1 { // checks point (0,last)
				fmt.Print("C")
			} else if row == y-1 && col == 0 { // checks point (last,0)
				fmt.Print("C")
			} else if row == y-1 && col == x-1 { // if (last,last)
				fmt.Print("A")
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 { // in between print 'B'
				fmt.Print("B")
			} else {
				fmt.Print(" ") // in the center nothing
			}
		}
		// Only print newline if it's not the last row
		if row != y-1 {
			fmt.Println()
		}
	}
}
