package main

import (
	"fmt"
)

func main() {
	printDiamond(2)
}

func printDiamond(n int) {
	fmt.Println("Мой бриллиант:")
	n = n - 1
	left := n
	right := n
	for i := 0; i <= n*2; i++ {
		for j := 0; j <= n*2; j++ {
			if j == left || j == right {
				fmt.Print("#")
			} else if j < right {
				fmt.Print("+")
			}
		}
		if i < n {
			left--
			right++
		} else {
			left++
			right--
		}

		fmt.Println()
	}
}
