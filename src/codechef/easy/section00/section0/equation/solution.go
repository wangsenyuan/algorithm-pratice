package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		var N, A, B, C int
		fmt.Scanf("%d %d %d %d", &N, &A, &B, &C)
		res := f(N) - f(N - A - 1) - f(N - B - 1) - f(N - C - 1) + f(N - A - B - 2) +
			f(N - A - C - 2) + f(N - B - C - 2) - f(N - A - B - C - 3)
		fmt.Println(res)
	}
}

func f(N int) int64 {
	M := int64(N)
	if M < 0 {
		return 0
	}

	return (M + 1) * (M + 2) * (M + 3) / 6
}
