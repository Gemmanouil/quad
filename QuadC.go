package main

import "fmt"

func QuadC(x, y int) { //check ama einai 8etikos alliws exit
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ { //loop seiras
		for j := 0; j < x; j++ { //loop sthlhs
			if i == 0 {
				if j == 0 || j == x-1 {
					fmt.Print("A")
				} else {
					fmt.Print("B")
				}
			} else if i == y-1 {
				if j == 0 || j == x-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			} else {
				if j == 0 || j == x-1 {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
		}
		//gia na mhn vgainoun sthn idia seira
		fmt.Print("\n")
	}
}

//paradeigma
func main() {
	QuadC(5, 3)
}
