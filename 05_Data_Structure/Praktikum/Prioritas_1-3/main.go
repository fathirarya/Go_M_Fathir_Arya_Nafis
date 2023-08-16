package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	unique := map[int]int{}
	result := []int{}

	for _, v := range angka {
		value, _ := strconv.Atoi(string(v))
		unique[value] += 1
	}

	for k, v := range unique {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
