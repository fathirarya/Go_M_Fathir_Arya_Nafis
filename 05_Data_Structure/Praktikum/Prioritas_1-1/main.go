package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayMerged := append(arrayA, arrayB...)
	values := map[string]bool{}
	result := []string{}

	for _, v := range arrayMerged {
		if !values[v] {
			values[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{"devil jin"}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
