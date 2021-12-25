package main

import (
	"fmt"
	"math"
)

func main() {
	var tc int

	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--

		var s string
		fmt.Scanf("%s", &s)

		fmt.Println(solve(s))
	}
}

var ii []int

func init() {
	for i := 1; 4*i+1 < 100000; i++ {
		x := 4*i + 1
		y := int(math.Sqrt(float64(x)))
		if y*y == x {
			ii = append(ii, i)
		}
	}
}

func solve(s string) int {
	n := len(s)
	cnt := make([]int, n+1)

	var res int

	for i := 0; i < n; i++ {
		cnt[i+1] = cnt[i]
		if s[i] == '1' {
			cnt[i+1]++
		}

		for k := 0; ii[k] <= i+1; k++ {
			l := ii[k]
			j := i + 1 - l
			a := cnt[i+1] - cnt[j]
			if a+a*a == l {
				res++
			}
		}
	}

	return res
}
