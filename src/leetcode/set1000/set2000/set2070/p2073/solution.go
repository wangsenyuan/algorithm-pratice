package p2073

func checkAlmostEquivalent(word1 string, word2 string) bool {
	x := freq(word1)
	y := freq(word2)
	for i := 0; i < 26; i++ {
		if abs(x[i]-y[i]) > 3 {
			return false
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func freq(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		res[int(s[i]-'a')]++
	}
	return res
}
