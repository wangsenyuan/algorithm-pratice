package main

import "fmt"

const MOD = 1000000007

func findPaths(m int, n int, N int, i int, j int) int {
	f := make([][][]int64, 2)

	for k := 0; k < 2; k++ {
		f[k] = make([][]int64, m)
		for i := 0; i < m; i++ {
			f[k][i] = make([]int64, n)
		}

	}

	p := 0
	f[p][i][j] = 1

	dp := func(f [][]int64, i, j int) int64 {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		return f[i][j]
	}

	var res int64

	for k := 0; k < N; k++ {
		q := 1 - p
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 {
					res = (res + f[p][i][j]) % MOD
				}
				if j == 0 {
					res = (res + f[p][i][j]) % MOD
				}
				if i == m-1 {
					res = (res + f[p][i][j]) % MOD
				}

				if j == n-1 {
					res = (res + f[p][i][j]) % MOD
				}

				tmp := int64(0)
				tmp = (tmp + dp(f[p], i-1, j)) % MOD
				tmp = (tmp + dp(f[p], i+1, j)) % MOD
				tmp = (tmp + dp(f[p], i, j-1)) % MOD
				tmp = (tmp + dp(f[p], i, j+1)) % MOD
				f[q][i][j] = tmp
			}
		}
		p = q
	}

	return int(res)
}

func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
	fmt.Println(findPaths(1, 3, 3, 0, 1))

}
