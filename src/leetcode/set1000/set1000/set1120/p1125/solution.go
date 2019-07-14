package p1125

const INF = 1 << 30

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n := len(req_skills)
	m := len(people)
	N := (1 << uint(n))

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = INF
		}
	}

	pp := make([]int, N)
	qq := make([]int, N)
	for j := 0; j < m; j++ {
		for mask := 0; mask < N; mask++ {
			dp[mask][j+1] = min(dp[mask][j+1], dp[mask][j])
		}

		dp[0][j] = 0

		ss := make(map[string]bool)
		for i := 0; i < len(people[j]); i++ {
			ss[people[j][i]] = true
		}

		var bit int

		for i := 0; i < n; i++ {
			if ss[req_skills[i]] {
				bit |= 1 << uint(i)
			}
		}

		for mask := 0; mask < N; mask++ {
			cur := dp[mask][j]
			if cur >= INF {
				continue
			}

			next := mask | bit
			if dp[next][j+1] > cur+1 {
				dp[next][j+1] = cur + 1
				pp[next] = mask
				qq[next] = j
			}
		}
	}

	var res []int

	for mask := N - 1; mask > 0; {
		res = append(res, qq[mask])
		mask = pp[mask]
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
