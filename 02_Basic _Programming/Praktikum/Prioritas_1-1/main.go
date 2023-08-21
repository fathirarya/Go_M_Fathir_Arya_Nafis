package main

import "fmt"

func main() {
	var a, b, h, luas float32
	fmt.Println("=====Luas Trapesium=====")
	fmt.Print("Masukkan Alas Atas	: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan Alas Bawah	: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan Tinggi		: ")
	fmt.Scan(&h)

	luas = (a + b) * h / 2
	fmt.Println("Luas Trapesium:", luas, "cm.")
}
