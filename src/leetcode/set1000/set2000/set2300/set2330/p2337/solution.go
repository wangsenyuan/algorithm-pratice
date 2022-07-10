package p2337

func canChange(start string, target string) bool {
	n := len(target)
	// 如果 target[i] == L
	// 则要么start[i] = L or start[i] = _ and some j > i, start[j] = L, and [i..j] is blank
	// 如果 target[i] = R
	// 相似
	var cnt int
	for i := 0; i < n; i++ {
		if target[i] == 'L' {
			cnt++
		}

		if start[i] == 'L' {
			cnt--
		}

		if cnt < 0 {
			return false
		}

		if cnt > 0 && start[i] == 'R' {
			return false
		}

		if target[i] == 'R' && cnt > 0 {
			return false
		}
	}

	if cnt != 0 {
		return false
	}

	for i := n - 1; i >= 0; i-- {
		if target[i] == 'R' {
			cnt++
		}
		if start[i] == 'R' {
			cnt--
		}
		if cnt < 0 {
			return false
		}

		if start[i] == 'L' && cnt > 0 {
			return false
		}
		if target[i] == 'L' && cnt > 0 {
			return false
		}
	}

	return cnt == 0
}
