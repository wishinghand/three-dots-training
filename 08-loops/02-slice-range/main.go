package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(numbers ...int) int {
	var numsum int
	for _, n := range numbers {
		numsum += n
	}
	return numsum	
}