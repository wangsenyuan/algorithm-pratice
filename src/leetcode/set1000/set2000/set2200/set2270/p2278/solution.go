package p2278

func percentageLetter(s string, letter byte) int {
	n := len(s)
	var cnt int

	for i := 0; i < n; i++ {
		if s[i] == letter {
			cnt++
		}
	}
	return (cnt * 100 / n)
}
