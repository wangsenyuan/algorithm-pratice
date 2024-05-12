package p3145

func findPermutationDifference(s string, t string) int {
	pos := make([]int, 26)
	for i := 0; i < len(s); i++ {
		pos[int(s[i]-'a')] = i
	}

	var res int

	for i := 0; i < len(t); i++ {
		x := int(t[i] - 'a')
		res += abs(i - pos[x])
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
