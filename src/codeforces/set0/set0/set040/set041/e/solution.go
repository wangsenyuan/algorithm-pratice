package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := solve(n)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
	}
	buf.WriteTo(os.Stdout)
}

func solve(n int) [][]int {
	var res [][]int

	if n < 6 {
		for i := 2; i <= n; i++ {
			res = append(res, []int{i - 1, i})
			if i-3 >= 1 {
				res = append(res, []int{i - 3, i})
			}
		}
	} else {
		for i := 1; i <= n/2; i++ {
			for j := n/2 + 1; j <= n; j++ {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
