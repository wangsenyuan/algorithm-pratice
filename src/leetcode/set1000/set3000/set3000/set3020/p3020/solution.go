package p3020

func flowerGame(n int, m int) int64 {
	// x + y is odd 吗？
	a := int64(n)
	b := int64(m)
	// 偶数+奇数
	res := a / 2 * (b - b/2)
	// 奇数+偶数
	res += (a - a/2) * (b / 2)

	return res
}
