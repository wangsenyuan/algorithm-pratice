package p1125

const INF = 1 << 30

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n := len(req_skills)
	m := len(people)
	N := (1 << uint(n))

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	pp := make([]int, N)
	qq := make([]int, N)
	for j := 0; j < m; j++ {

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
			cur := dp[mask]
			if cur >= INF {
				continue
			}

			next := mask | bit
			if dp[next] > cur+1 {
				dp[next] = cur + 1
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
