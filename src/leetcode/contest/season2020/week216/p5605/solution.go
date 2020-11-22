package p5605

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var x, y int
	var i, j int

	for x < len(word1) && y < len(word2) {
		if word1[x][i] != word2[y][j] {
			return false
		}
		i++
		if i == len(word1[x]) {
			i = 0
			x++
		}
		j++
		if j == len(word2[y]) {
			j = 0
			y++
		}
	}
	return x == len(word1) && y == len(word2)
}
