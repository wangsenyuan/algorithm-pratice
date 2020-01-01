package problem4

import "math/bits"

func domino(n int, m int, broken [][]int) int {
	M := 1 << uint(m)

	mp := make([]int, n+1)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, M)
	}

	for _, brk := range broken {
		a, b := brk[0], brk[1]
		mp[a] |= 1 << uint(b)
	}

	mp[n] = M - 1

	var ans int

	for i := 0; i < n; i++ {
		for state := 0; state < M; state++ {
			if state&mp[i] != mp[i] {
				continue
			}
			if i == n-1 {
				update(&ans, dp[i][state])
			}
			for j := 0; j < m-1; j++ {
				if state&(1<<uint(j)) == 0 && state&(1<<uint(j+1)) == 0 {
					update(&dp[i][state|(1<<uint(j))|(1<<uint(j+1))], dp[i][state]+1)
				}
			}
		}

		for state := 0; state < M; state++ {
			if state&mp[i] != mp[i] {
				continue
			}

			for next := 0; next < M; next++ {
				if state&next == 0 && mp[i+1]&next == 0 {
					// vertical put from the positions that next is set
					update(&dp[i+1][mp[i+1]|next], dp[i][state]+bits.OnesCount32(uint32(next)))
				}
			}
		}
	}

	return ans
}

func update(a *int, b int) {
	if *a < b {
		*a = b
	}
}
