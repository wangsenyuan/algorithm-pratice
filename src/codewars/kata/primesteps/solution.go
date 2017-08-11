package main

import "fmt"

func Step(g, m, n int) []int {
	nums := make([]bool, n+1)

	primes := make(map[int]bool)

	x := 2

	for x <= n {
		if !nums[x] {
			if x >= m {
				primes[x] = true
			}

			a := x - g

			if _, found := primes[a]; found {
				return []int{a, x}
			}

			for y := 2 * x; y <= n; y += x {
				nums[y] = true
			}
		}
		x++
	}

	return nil
}

func main() {
	fmt.Println(Step(2, 100, 110))
	fmt.Println(Step(4, 100, 110))
	fmt.Println(Step(6, 100, 110))
	fmt.Println(Step(8, 300, 400))
	fmt.Println(Step(10, 300, 400))
	fmt.Println(Step(11, 30000, 100000))
}
