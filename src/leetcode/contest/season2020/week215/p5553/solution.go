package p5553

func canDistribute1(nums []int, quantity []int) bool {
	//sort.Ints(nums)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	m := len(quantity)
	M := 1 << m

	xs := make([]int, 0, len(freq))
	for _, v := range freq {
		xs = append(xs, v)
	}
	n := len(xs)
	dp := make([]bool, M)
	fp := make([]bool, M)
	dp[0] = true

	for i := 0; i < n; i++ {
		for state := 0; state < M; state++ {
			if dp[state] {
				fp[state] = true
				continue
			}
			for j := state; j > 0; j = (j - 1) & state {
				prev := state - j
				if !dp[prev] {
					continue
				}
				var tot int
				for k := 0; k < m; k++ {
					if j&(1<<k) > 0 {
						tot += quantity[k]
					}
				}
				if tot <= xs[i] {
					fp[state] = true
					break
				}
			}
		}
		for state := 0; state < M; state++ {
			dp[state] = fp[state]
			fp[state] = false
		}
	}

	return dp[M-1]
}

func canDistribute(nums []int, quantity []int) bool {
	//sort.Ints(nums)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	m := len(quantity)
	M := 1 << m

	dp := make([]bool, M)
	fp := make([]bool, M)
	dp[0] = true

	for _, v := range freq {
		for state := 0; state < M; state++ {
			if !dp[state] {
				continue
			}
			diff := (M - 1) ^ state
			for j := diff; j > 0; j = (j - 1) & diff {
				//if state&j > 0 {
				//	continue
				//}
				var tot int
				for k := 0; k < m; k++ {
					if j&(1<<k) > 0 {
						tot += quantity[k]
					}
				}
				if tot <= v {
					fp[state|j] = true
				}
			}
		}
		for state := 0; state < M; state++ {
			if !dp[state] {
				dp[state] = fp[state]
			}
			fp[state] = false
		}
	}

	return dp[M-1]
}
