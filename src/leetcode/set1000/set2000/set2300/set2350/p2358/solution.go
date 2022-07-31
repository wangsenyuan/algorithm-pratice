package p2358

func maximumGroups(grades []int) int {
	// 1 + 2 + 3 + 4.. <= n
	n := len(grades)
	// (1 + i) * i / 2 <= n
	i := 1
	for (1 + i) < 2*n/i {
		i++
	}
	if (1+i)*i > 2*n {
		i--
	}
	return i
}
