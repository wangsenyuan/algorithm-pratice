package p1093

func sampleStats(count []int) []float64 {
	var n int

	for i := 0; i < len(count); i++ {
		n += count[i]
	}

	m1 := n / 2
	m2 := m1
	if n&1 == 0 {
		m2++
	}
	mi := 255
	ma := 0
	var mean float64
	var median float64
	var cnt int
	var mode int
	var m_cnt int
	for i := 0; i < len(count); i++ {
		if count[i] > 0 {
			mi = min(mi, i)
			ma = max(ma, i)
		}

		mean += float64(i) * float64(count[i]) / float64(n)

		if cnt < m1 && cnt+count[i] >= m1 {
			median += float64(i) * 0.5
		}
		if cnt < m2 && cnt+count[i] >= m2 {
			median += float64(i) * 0.5
		}

		if m_cnt < count[i] {
			mode = i
		}
		m_cnt = max(m_cnt, count[i])

		cnt += count[i]
	}

	return []float64{float64(mi), float64(ma), mean, median, float64(mode)}
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
