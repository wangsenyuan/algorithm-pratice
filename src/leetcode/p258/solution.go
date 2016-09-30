package main

import "fmt"

func main() {
	fmt.Println(addDigits(65536))
}

func addDigits(n int) int {
	return n - 9*((n-1)/9)
}
