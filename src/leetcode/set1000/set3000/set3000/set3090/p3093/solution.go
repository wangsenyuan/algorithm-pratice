package p3093

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	next := make([][]int, 0, 1)
	var ans []int

	addNode := func() {
		next = append(next, make([]int, 26))
		ans = append(ans, -1)
	}
	//root
	addNode()

	for i, word := range wordsContainer {
		s := reverse(word)
		var node int
		for j := 0; j < len(s); j++ {
			if ans[node] < 0 || len(s) < len(wordsContainer[ans[node]]) {
				ans[node] = i
			}
			x := int(s[j] - 'a')
			if next[node][x] == 0 {
				addNode()
				next[node][x] = len(next) - 1
			}
			y := next[node][x]

			node = y
		}
		if ans[node] < 0 || len(s) < len(wordsContainer[ans[node]]) {
			ans[node] = i
		}
	}

	query := func(s string) int {
		var node int
		for i := 0; i < len(s); i++ {
			x := int(s[i] - 'a')
			if next[node][x] == 0 {
				return ans[node]
			}
			node = next[node][x]
		}
		return ans[node]
	}

	res := make([]int, len(wordsQuery))

	for i, s := range wordsQuery {
		res[i] = query(reverse(s))
	}

	return res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
