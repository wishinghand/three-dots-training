package main

import "fmt"

func Alphabet(length int) []string {
	var alphabet []string
	for x:= 0; x < length; x++ {
		alphabet = append(alphabet, characterByIndex(x))
	}
	return alphabet
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
