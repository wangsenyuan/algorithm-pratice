package p806

func numberOfLines(widths []int, S string) []int {
	n := len(S)
	if n == 0 {
		return []int{0, 0}
	}
	var lines int
	var width int
	for i := 0; i < n; i++ {
		c := S[i]
		x := int(c - 'a')
		if width+widths[x] > 100 {
			lines++
			width = widths[x]
		} else {
			width += widths[x]
		}
	}

	return []int{lines + 1, width}
}
