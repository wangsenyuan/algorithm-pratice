package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var N, M, A int
		fmt.Fscanf(f, "%d %d %d", &N, &M, &A)
		res := solve(N, M, A)
		if res == nil {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
		} else {
			fmt.Printf("Case #%d: %d %d %d %d %d %d\n", i, res[0], res[1], res[2], res[3], res[4], res[5])
		}
	}
}

func solve(N, M, A int) []int {
	if N*M < A {
		return nil
	}

	k := A / M

	if k*M == A {
		return []int{0, 0, k, 0, k, M}
	}

	// k * M <= A & (k +1) * M > A

	return []int{0, A - k*M, k, 0, k + 1, M}
}
