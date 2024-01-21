package p3003

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e && (c != e || ok(b, d, f)) ||
		b == f && (d != f || ok(a, c, e)) ||
		c+d == e+f && (a+b != e+f || ok(c, a, e)) ||
		c-d == e-f && (a-b != e-f || ok(c, a, e)) {
		return 1
	}
	return 2
}

func ok(l, m, r int) bool {
	return m < min(l, r) || m > max(l, r)
}
