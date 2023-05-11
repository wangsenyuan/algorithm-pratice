package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	var tc int
	fmt.Scanf("%d", &tc)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var m, n int
		var x float64
		fmt.Scanf("%d%d%f", &m, &n, &x)
		res := solve(m, n, x)
		buf.WriteString(fmt.Sprintf("%.10f\n", res))
	}
	fmt.Print(buf.String())
}

func solve(m int, n int, x float64) float64 {

	used := make([]int, 51)

	f := math.Cos
	pow := math.Pow

	comb := func(n int, k int) float64 {
		res := 1.0
		for i := 0; i < k; i++ {
			res *= float64(n - i)
			res /= float64(i + 1)
		}
		return res
	}

	var dfs func(now int, rest int) float64

	dfs = func(now int, rest int) float64 {
		if now == 1 {
			used[1] = rest
			var res float64 = 1
			cnt := m
			for i := n; i > 0; i-- {
				res *= comb(cnt, used[i])
				cnt -= used[i]
				res *= pow(f(float64(i)*x/float64(n)), float64(used[i]))
			}
			return res
		}
		var res float64
		for i := 0; now*i <= rest; i++ {
			used[now] = i
			res += dfs(now-1, rest-now*i)
		}
		return res
	}

	return dfs(n, n)
}
