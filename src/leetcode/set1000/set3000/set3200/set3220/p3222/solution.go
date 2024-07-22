package p3222

func losingPlayer(x int, y int) string {
	// x个75， y个10
	// 组成 115
	// 15, 2
	// 23 = 15 + 2 * 4 (没有其他的组合)
	// 也就是一个x，消耗4个y
	if y > 4*x {
		y = 4 * x
	}
	if x*4 > y {
		x = y / 4
	}
	if x%2 == 1 {
		return "Alice"
	}
	return "Bob"
}
