package p2007

import "sort"

func findOriginalArray(changed []int) []int {
	n := len(changed)
	if n%2 == 1 {
		return nil
	}
	origin := make([]int, 0, n/2)
	sort.Ints(changed)
	// A[j] = 2 * A[i]
	j := 1
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		j = max(j, i+1)
		origin = append(origin, changed[i])
		// we need to find some j, that A[j] = 2 * A[i]
		for j < n && (used[j] || 2*changed[i] >= changed[j]) {
			if !used[j] && 2*changed[i] == changed[j] {
				break
			}
			j++
		}
		if j == n || 2*changed[i] != changed[j] {
			return nil
		}
		used[j] = true
		j++
	}

	if len(origin) == n/2 {
		return origin
	}

	return nil
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
