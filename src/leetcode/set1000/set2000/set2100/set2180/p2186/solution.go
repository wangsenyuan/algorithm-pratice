package p2186

func minSteps(s string, t string) int {
	f1 := freq(s)
	f2 := freq(t)

	var res int

	for i := 0; i < 26; i++ {
		x := max(f1[i], f2[i])
		res += x - f1[i] + x - f2[i]
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func freq(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		res[int(s[i]-'a')]++
	}
	return res
}
