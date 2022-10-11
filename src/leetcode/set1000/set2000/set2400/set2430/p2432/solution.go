package p2432

func hardestWorker(n int, logs [][]int) int {
	var best int
	var ans int
	var cur int
	for _, log := range logs {
		j := log[0]
		leave := log[1]
		took := leave - cur
		if took > best || took == best && j < ans {
			best = took
			ans = j
		}
		cur = leave
	}

	return ans
}
