package main

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return k
	}
	same, diff := k, k*(k-1)
	for i := 2; i < n; i++ {
		same, diff = diff, (k-1)*(same+diff)
	}

	return same + diff
}
