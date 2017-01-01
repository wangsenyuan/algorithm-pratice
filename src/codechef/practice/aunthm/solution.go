package main

import "fmt"

func main() {
	var t, n, m int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		fmt.Scanf("%d %d", &n, &m)

		fmt.Printf("%.6f\n", float64(n + m - 1))
	}
}
