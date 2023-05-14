package p2675

func circularGameLosers(n int, k int) []int {
	d := k
	vis := make([]bool, n+1)
	vis[0] = true
	it := 0
	for {
		it = (it + d) % n
		if vis[it] {
			break
		}
		vis[it] = true
		d += k
	}

	var res []int
	for i := 0; i < n; i++ {
		if !vis[i] {
			res = append(res, i+1)
		}
	}

	return res
}
