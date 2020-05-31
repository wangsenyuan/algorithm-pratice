package p1467

import (
	"fmt"
	"math/big"
	"os"
)

const MAX_X = 60

var F []*big.Float

func init() {
	F = make([]*big.Float, MAX_X)

	F[0] = big.NewFloat(1)
	for i := 1; i < MAX_X; i++ {
		F[i] = big.NewFloat(1)
		F[i].Mul(F[i-1], big.NewFloat(float64(i)))
		// fmt.Fprintf(os.Stderr, "%e\n", F[i])
	}

}

func getProbability(balls []int) float64 {

	n := len(balls)
	a := make([]int, n)
	b := make([]int, n)

	sum := func(arr []int) int {
		var res int

		for i := 0; i < len(arr); i++ {
			res += arr[i]
		}

		return res
	}

	color := func(arr []int) int {
		var a int

		for i := 0; i < len(arr); i++ {
			if arr[i] > 0 {
				a++
			}

		}

		return a
	}

	total := sum(balls)
	all := big.NewFloat(0)
	res := big.NewFloat(0)
	var dfs func(int)

	dfs = func(i int) {
		if i == n {
			m := sum(a)
			if m*2 == total {
				x := color(a)
				y := color(b)
				cnt := big.NewFloat(1)
				cnt.Mul(F[m], F[m])
				cnt.Quo(cnt, comb(a))
				cnt.Quo(cnt, comb(b))
				if x == y {
					res.Add(res, cnt)
				}
				all.Add(all, cnt)

			}
			return
		}

		for x := 0; x <= balls[i]; x++ {
			a[i] = x
			b[i] = balls[i] - x
			dfs(i + 1)
		}
	}

	dfs(0)

	fmt.Fprintf(os.Stderr, "%e %e\n", res, all)

	res.Quo(res, all)

	ans, _ := res.Float64()
	return ans
}

func comb(arr []int) *big.Float {
	var res = big.NewFloat(1)

	for i := 0; i < len(arr); i++ {
		res.Mul(res, F[arr[i]])
	}

	return res
}
