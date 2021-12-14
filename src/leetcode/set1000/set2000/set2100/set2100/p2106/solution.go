package p2106

import "sort"

func maxTotalFruits1(fruits [][]int, startPos int, k int) int {
	// n <= 10^^5
	// k <= 10 ^^ 5
	n := len(fruits)
	p := sort.Search(n, func(i int) bool {
		return fruits[i][0] > startPos
	})

	sum := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sum[i] = fruits[i][1]
		if i+1 < n {
			sum[i] += sum[i+1]
		}
	}
	var best int
	// go left
	for i := p; i >= 0; i-- {
		// if go right as far as possible, then go left, and stop at i
		if i == n || fruits[i][0] > startPos {
			continue
		}
		xi := fruits[i][0]
		if startPos-xi > k {
			break
		}
		tmp := sum[i]
		if p < n {
			best = max(best, tmp-sum[p])
		} else {
			best = max(best, tmp)
		}
		// go directly to xi
		// x - xi + x - startPos > k
		j := sort.Search(n, func(j int) bool {
			x := fruits[j][0]
			return x-xi+x-startPos > k
		})
		if j < n {
			tmp -= sum[j]
		}
		best = max(best, tmp)
	}
	for i := 0; i < n; i++ {
		sum[i] = fruits[i][1]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	if p > 0 && fruits[p-1][0] == startPos {
		p--
	}

	for i := p; i < n; i++ {
		if fruits[i][0] < startPos {
			continue
		}
		xi := fruits[i][0]
		if xi-startPos > k {
			break
		}
		tmp := sum[i]
		if p > 0 {
			best = max(best, tmp-sum[p-1])
		} else {
			best = max(best, tmp)
		}
		// xi - x + startPos - x > k
		j := sort.Search(i, func(j int) bool {
			x := fruits[j][0]
			return xi-x+startPos-x <= k
		})

		if j > 0 {
			tmp -= sum[j-1]
		}
		best = max(best, tmp)
	}

	return best
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	// n <= 10^^5
	// k <= 10 ^^ 5
	n := len(fruits)
	p := sort.Search(n, func(i int) bool {
		return fruits[i][0] > startPos
	})

	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = fruits[i][1] + sum[i+1]
	}
	var best int
	// go left
	for i, j := p, n; i >= 0; i-- {
		// if go right as far as possible, then go left, and stop at i
		if i == n || fruits[i][0] > startPos {
			continue
		}

		xi := fruits[i][0]

		if startPos-xi > k {
			break
		}

		best = max(best, sum[i]-sum[p])

		for j > p && (fruits[j-1][0]-startPos+fruits[j-1][0]-xi > k) {
			j--
		}

		if j < n && fruits[j][0] < startPos {
			continue
		}

		best = max(best, sum[i]-sum[j])
	}
	sum[0] = 0
	for i := 0; i < n; i++ {
		sum[i+1] = fruits[i][1] + sum[i]
	}

	if p > 0 && fruits[p-1][0] == startPos {
		p--
	}

	for i, j := p, 0; i < n; i++ {
		xi := fruits[i][0]
		if xi-startPos > k {
			break
		}

		best = max(best, sum[i+1]-sum[p])

		for j < p && startPos-fruits[j][0]+xi-fruits[j][0] > k {
			j++
		}

		if fruits[j][0] > startPos {
			continue
		}

		best = max(best, sum[i+1]-sum[j])
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
