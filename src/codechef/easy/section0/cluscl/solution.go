package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var p, q, n int64
		fmt.Scanf("%d %d %d", &p, &q, &n)
		n--
		den := 2 * abs(p-q)
		num := q * n

		ans := num/den + 1

		fmt.Println(ans)

		t--
	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
