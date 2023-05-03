package lcp66

func minNumBooths(demand []string) int {
	cnt := make([]int, 26)
	tmp := make([]int, 26)
	for _, de := range demand {
		for i := 0; i < len(de); i++ {
			x := int(de[i] - 'a')
			tmp[x]++
		}
		for i := 0; i < 26; i++ {
			cnt[i] = max(cnt[i], tmp[i])
			tmp[i] = 0
		}
	}
	var res int
	for _, v := range cnt {
		res += v
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
