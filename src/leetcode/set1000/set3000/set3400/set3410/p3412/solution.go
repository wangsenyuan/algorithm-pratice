package p3412

func calculateScore(s string) int64 {
	var score int
	pos := make([][]int, 26)

	n := len(s)
	for i := range n {
		x := int(s[i] - 'a')
		y := 25 - x
		if len(pos[y]) > 0 {
			score += i - pos[y][len(pos[y])-1]
			pos[y] = pos[y][:len(pos[y])-1]
		} else {
			pos[x] = append(pos[x], i)
		}
	}
	return int64(score)
}
