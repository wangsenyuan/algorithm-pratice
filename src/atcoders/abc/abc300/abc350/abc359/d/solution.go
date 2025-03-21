package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)

	res := solve(s[:n], k)
	fmt.Println(res)
}

const mod = 998244353

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

func solve(s string, k int) int {
	n := len(s)
	K := 1 << k
	dp := make([]int, K)
	var valid []int
	// 初始化
	for state := 0; state < K; state++ {
		if checkPalindrome(state, k) {
			continue
		}
		valid = append(valid, state)
		ok := true
		for i := 0; i < k; i++ {
			x := (state >> (k - 1 - i)) & 1
			if x == 0 && s[i] == 'B' || x == 1 && s[i] == 'A' {
				ok = false
				break
			}
		}
		if ok {
			dp[state] = 1
		}
	}

	get := func(i int) int {
		if s[i] == 'A' {
			return 0
		}
		if s[i] == 'B' {
			return 1
		}
		return 2
	}

	xs := [][]int{{0}, {1}, {0, 1}}

	mask := K - 1
	fp := make([]int, K)
	for i := k; i < n; i++ {
		clear(fp)
		j := get(i)
		for _, x := range xs[j] {
			for _, state := range valid {
				if dp[state] == 0 {
					// conflicts
					continue
				}
				next := ((state << 1) | x) & mask
				fp[next] = add(fp[next], dp[state])
			}
		}

		copy(dp, fp)
	}

	var res int
	for _, state := range valid {
		res = add(res, dp[state])
	}
	return res
}

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

func checkPalindrome(state int, k int) bool {
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		x := (state >> i) & 1
		y := (state >> j) & 1
		if x != y {
			return false
		}
	}
	return true
}
