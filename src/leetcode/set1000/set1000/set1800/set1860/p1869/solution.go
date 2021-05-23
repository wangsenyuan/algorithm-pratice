package p1869

func checkZeroOnes(s string) bool {
	a := findLongest(s, '1')
	b := findLongest(s, '0')
	return a > b
}

func findLongest(s string, x byte) int {
	var best int
	for i := 0; i < len(s); {
		if s[i] != x {
			i++
			continue
		}
		var cnt int
		for i < len(s) && s[i] == x {
			i++
			cnt++
		}
		if best < cnt {
			best = cnt
		}
	}
	return best
}
