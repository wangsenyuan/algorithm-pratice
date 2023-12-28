package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)

	for tc > 0 {
		tc--
		var n, k int
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		res := solve(a, k)
		fmt.Fprintln(out, res)
	}

}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

const H = 7

func solve(a []int, k int) int {
	// 如果 a[l...r] and的有m个bit set
	// a[l...r+1] <= m
	// two pointers
	// dp[state] = 以state为and的结果
	T := 1 << H
	dp := make([]int, T)
	dp[T-1] = 1

	fp := make([]int, T)

	for i := 0; i < len(a); i++ {
		for state := 0; state < T; state++ {
			fp[state&a[i]] = add(fp[state&a[i]], dp[state])
		}
		for state := 0; state < T; state++ {
			// take or not take
			dp[state] = add(dp[state], fp[state])
			fp[state] = 0
		}
	}
	var res int

	for i := 0; i < T; i++ {
		if digitCount(i) == k {
			res = add(res, dp[i])
		}
	}

	return res
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
