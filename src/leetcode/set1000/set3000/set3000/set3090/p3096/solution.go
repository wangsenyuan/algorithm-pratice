package p3096

func minimumLevels(possible []int) int {
	n := len(possible)

	suf := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if possible[i] == 0 {
			suf[i]--
		} else {
			suf[i]++
		}
	}

	var pref int
	for i := 0; i < n; i++ {
		if possible[i] == 0 {
			pref--
		} else {
			pref++
		}
		if i+1 < n && pref > suf[i+1] {
			return i + 1
		}
	}
	return -1
}
