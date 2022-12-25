package p2511

func captureForts(forts []int) int {

	var res int

	n := len(forts)

	pos := make([]int, 2)

	for i := 0; i < n; i++ {
		if forts[i] == 0 {
			// empty
			continue
		}
		if forts[i] == 1 {
			if pos[0] > pos[1] {
				res = max(res, (i+1)-pos[0]-1)
			}
			pos[1] = i + 1
		} else {
			// forst[i] == -1
			if pos[1] > pos[0] {
				res = max(res, (i+1)-pos[1]-1)
			}
			pos[0] = i + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
