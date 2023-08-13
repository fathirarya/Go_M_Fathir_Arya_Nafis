package main

import "fmt"

func main() {
	var baris int
	fmt.Print("Input: ")
	fmt.Scan(&baris)

	fmt.Println("Output: ")
	for i := 1; i <= baris; i++ {
		for j := 1; j <= baris-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
