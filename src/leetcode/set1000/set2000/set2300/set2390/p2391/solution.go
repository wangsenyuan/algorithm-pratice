package p2391

func garbageCollection(garbage []string, travel []int) int {
	n := len(garbage)

	cnt := make([][]int, n)

	for i := 0; i < n; i++ {
		cur := garbage[i]
		cnt[i] = make([]int, 3)
		for j := 0; j < len(cur); j++ {
			if cur[j] == 'G' {
				cnt[i][0]++
			} else if cur[j] == 'P' {
				cnt[i][1]++
			} else {
				cnt[i][2]++
			}
		}
	}
	var res int
	for x := 0; x < 3; x++ {
		var tmp int
		for i := 0; i < n; i++ {
			if cnt[i][x] > 0 {
				res += tmp + cnt[i][x]
				tmp = 0
			}
			if i+1 < n {
				tmp += travel[i]
			}
		}
	}

	return res
}
