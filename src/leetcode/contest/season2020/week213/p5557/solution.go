package p5557

import "bytes"

const N = 30

var C [N][N]int

func init() {
	C[0][0] = 1
	for i := 1; i < N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}
}

func nCr(n, r int) int {
	return C[n][r]
}

func kthSmallestPath(destination []int, k int) string {
	// at (i, j), there are C[r - i][c - j] strings after
	// if go right,  SH...
	//or go down , SV...
	r := destination[0]
	c := destination[1]
	var x, y int
	var buf bytes.Buffer
	for x != r || y != c {
		if x == r {
			// can only go right
			buf.WriteByte('H')
			y++
			continue
		}
		if y == c {
			buf.WriteByte('V')
			x++
			continue
		}

		dx := r - x
		dy := c - y
		tot := dx + dy

		if nCr(tot-1, dy-1) >= k {
			// go right
			buf.WriteByte('H')
			y++
		} else {
			buf.WriteByte('V')
			x++
			k -= nCr(tot-1, dy-1)
		}
	}
	return buf.String()
}
