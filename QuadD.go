package main

import "fmt"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 { //checks if x,y positive
		return
	}

	for i := 0; i < y; i++ { //loop for row
		for j := 0; j < x; j++ { //loop for column
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				if j == 0 {
					fmt.Print("A")
				} else {
					fmt.Print("C")
				}
			} else if i == 0 || i == y-1 {
				fmt.Print("B")
			} else if j == 0 || j == x-1 {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() //gia na mhn einai sthn idia seira
	}
}

//paradeigma
func main() {
	QuadD(5, 3)
}
