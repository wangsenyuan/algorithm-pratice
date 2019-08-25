package p1111

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, seq string, expect []int) {
	res := maxDepthAfterSplit(seq)

	x, y := split(seq, expect)
	a, b := split(seq, expect)

	ans1 := max(maxDepth(x), maxDepth(y))
	ans2 := max(maxDepth(a), maxDepth(b))
	if ans1 != ans2 {
		t.Errorf("Sample %s, expect %v, but got %v,", seq, expect, res)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func split(s string, index []int) (string, string) {
	var buf bytes.Buffer
	var buf2 bytes.Buffer
	for i := 0; i < len(s); i++ {
		if index[i] == 1 {
			buf.WriteByte(s[i])
		} else {
			buf2.WriteByte(s[i])
		}
	}
	return buf.String(), buf2.String()
}

func maxDepth(s string) int {
	var level int
	var ans int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			level++
			if level > ans {
				ans = level
			} else {
				level--
			}
		}
	}

	return ans
}

func TestSample1(t *testing.T) {
	seq := "(()())"
	expect := []int{0, 1, 1, 1, 1, 0}

	runSample(t, seq, expect)
}

func TestSample2(t *testing.T) {
	seq := "()(())()"
	expect := []int{0, 0, 0, 1, 1, 0, 1, 1}

	runSample(t, seq, expect)
}
