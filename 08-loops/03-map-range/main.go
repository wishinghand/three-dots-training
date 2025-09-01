package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(obj map[int]string) []int {
	var keys []int
	for key := range obj {
		keys = append(keys, key)
	}
	return keys
}

func Values(obj map[int]string) []string {
	var maps []string
	for _, mapp := range obj {
		maps = append(maps, mapp)
	}
	return maps
}