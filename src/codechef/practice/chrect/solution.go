package main

import "fmt"

func main() {
	var t, n, m, k int
	fmt.Scanf("%d", &t)

	for t > 0 {
		fmt.Scanf("%d %d %d", &n, &m, &k)
		ans := dropStones(n, m, k)

		fmt.Println(ans)

		t--
	}
}
func dropStones(n int, m int, k int) int {
	if n > m {
		n, m = m, n
	}

	if n == 1 && m <= 2 {
		return 0
	} else if n == 1 {
		return k
	} else {
		return (k + 1) / 2
	}
}
