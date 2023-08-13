package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&nilai)

	switch {
	case nilai < 0 || nilai > 100:
		fmt.Println("Nilai Invalid")
	case nilai >= 80:
		fmt.Println("Grade A")
	case nilai >= 65:
		fmt.Println("Grade B")
	case nilai >= 50:
		fmt.Println("Grade C")
	case nilai >= 35:
		fmt.Println("Grade D")
	default:
		fmt.Println("Grade E")

	}
}
