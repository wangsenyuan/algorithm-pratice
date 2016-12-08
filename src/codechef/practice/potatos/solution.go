package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		x, y := 0, 0

		fmt.Scanf("%d %d\n", &x, &y)

		z := nextPrime(x + y)

		fmt.Println(z - x - y)
		t--
	}
}

func nextPrime(a int) int {
	nonPrimes := make([]bool, 3001)

	i := 2
	for i < len(nonPrimes) {
		for j := 2; i*j < len(nonPrimes); j++ {
			nonPrimes[i*j] = true
		}

		k := i + 1
		for nonPrimes[k] {
			k++
		}
		i = k
		if i > a {
			return i
		}
	}

	return -1
}
