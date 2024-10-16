package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	ans, swaps := solve(s)

	if expect != ans {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}

	l, r := swaps[0], swaps[1]
	buf := []byte(s)
	buf[l-1], buf[r-1] = buf[r-1], buf[l-1]

	var best int
	n := len(buf)
	for i := 0; i < n; i++ {
		var cnt int
		var level int

		for j := 0; j < n; j++ {
			if buf[(i+j)%n] == '(' {
				level++
			} else {
				level--
			}
			if level < 0 {
				cnt = -1
				break
			}
			if level == 0 {
				cnt++
			}
		}
		if level == 0 {
			best = max(best, cnt)
		}
	}

	if best != expect {
		t.Fatalf("Sample result %d %v, not correct", ans, swaps)
	}
}

func TestSample1(t *testing.T) {
	s := "()()())(()"
	expect := 5
	runSample(t, s, expect)
}
