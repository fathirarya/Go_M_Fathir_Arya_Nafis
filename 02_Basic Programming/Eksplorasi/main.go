package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	length := len(input)

	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kata := scanner.Text()
	fmt.Println("Captured:", kata)

	if isPalindrome(kata) {
		fmt.Println(kata, "adalah palindrome.")
	} else {
		fmt.Println(kata, "bukan palindrome.")
	}
}
