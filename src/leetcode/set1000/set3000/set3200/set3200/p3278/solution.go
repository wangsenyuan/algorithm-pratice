package p3278

func numberOfAlternatingGroups(colors []int) int {
	var res int
	n := len(colors)
	for i := 0; i < n; i++ {
		j := (i - 2 + n) % n
		k := (i - 1 + n) % n
		if colors[j] == colors[i] && colors[k] != colors[i] {
			res++
		}
	}
	return res
}
