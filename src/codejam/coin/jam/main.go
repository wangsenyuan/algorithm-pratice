package main

import "fmt"

func convertBinaryToBase(x int64, base int) int64 {
	if x == 0 {
		return 0
	}

	return int64(base)*convertBinaryToBase(x/2, base) + (x % 2)
}

func findFactor(x int64) (int64, bool) {
	for i := int64(2); i*i <= x; i++ {
		if x%i == 0 {
			return i, true
		}
	}
	return -1, false
}

func process(t int, n int, x int) {
	fmt.Printf("Case #%d:\n", t)
	for i := int64(1<<(uint(n)-1)) + 1; x > 0; i += 2 {
		factors := make([]int64, 0, 9)
		for j := 2; j <= 10; j++ {
			y := convertBinaryToBase(i, j)

			if factor, found := findFactor(y); found {
				factors = append(factors, factor)
			}
		}

		if len(factors) < 9 {
			continue
		}

		fmt.Printf("%d", convertBinaryToBase(i, 10))
		for _, factor := range factors {
			fmt.Printf(" %d", factor)
		}
		fmt.Println()
		x--
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var n, x int
		fmt.Scanf("%d %d", &n, &x)
		process(i, n, x)
	}
}
