package main

import "fmt"

func QuadB(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	for i := 0; i < y; i++ { // loop seiras
		for j := 0; j <= x; j++ { //loop stilis

			// Program #3 - tiponei mono tin timi (1,1)

			if x == 1 && y == 1 {
				fmt.Println("/")
				fmt.Println()
				return
			}

			// Program #2 - tiponei mono tin proti grammi

			if y == 1 && j == 0 {
				fmt.Print("/")
			} else if j == x-1 {
				fmt.Println("EDO")
				fmt.Println("\\")
				return
			} else if j != x-1 { // tiponei asterakia "*" mexri -1 theseis apo to telos tou "x"
				fmt.Print("*")
			}

			// Program #4 - tiponei mono tin Proti stili

			if y > 1 && x == 0 {
				fmt.Println(i)
				fmt.Println(j)
				fmt.Print("/")
			} else if i == y-1 {
				fmt.Println("\\")
				return
			} else if i != y-2 {
				fmt.Print("*")
			}
			//	if i != y-2 { // tiponei asterakia "*" mexri -2 theseis apo to telos tou "y"
			//		fmt.Print("*")
			//	}

		}
	}
	fmt.Println()
}

func main() {
	QuadB(1, 5)
}
