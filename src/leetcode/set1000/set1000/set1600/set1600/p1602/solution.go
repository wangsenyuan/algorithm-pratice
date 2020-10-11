package p1602

func maxDepth(s string) int {

	var level int
	var best int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			level++
		} else if s[i] == ')' {
			level--
		}
		if level > best {
			best = level
		}
	}
	return best
}
