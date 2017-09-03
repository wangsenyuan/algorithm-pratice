package main

import "fmt"

func flipLights(n int, m int) int {

	if m == 0 || n == 0 {
		return 1

	}

	if n == 1 {
		return 2
	}
	if n == 2 {
		if m == 1 {
			return 3
		}
		return 4
	}

	if m == 1 {
		return 4
	}

	if m == 2 {
		return 7
	}

	return 8
}

func main() {
	fmt.Println(flipLights(2, 1))
}
