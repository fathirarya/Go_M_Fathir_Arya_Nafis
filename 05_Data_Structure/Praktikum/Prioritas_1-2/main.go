package main

import "fmt"

func Mapping(slice []string) map[string]int {
	count := map[string]int{}

	for _, v := range slice {
		count[v] += 1
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
