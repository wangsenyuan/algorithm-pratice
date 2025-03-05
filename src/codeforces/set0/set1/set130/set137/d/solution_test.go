package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	k, x, ops, res := process(bufio.NewReader(strings.NewReader(s)))

	check := func(ans string) int {
		// 不能超过k段
		w := strings.Count(ans, "+")
		if w+1 > k {
			t.Fatalf("Sample result %s, not correct, it has %d substrings, > %d", ans, w, k)
		}
		ans = strings.ReplaceAll(ans, "+", "")
		var cnt int
		for i := 0; i < len(x); i++ {
			if x[i] != ans[i] {
				cnt++
			}
		}
		return cnt
	}

	u := check(expect)
	v := check(res)

	if u != v || v != ops {
		t.Fatalf("Sample result %d %s, not correct", ops, res)
	}
}

func TestSample1(t *testing.T) {
	s := `abacaba
1
`
	expect := "abacaba"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abdcaba
2
`
	expect := "abdcdba"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abdcaba
5
`
	expect := "a+b+d+c+aba"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `abacababababbcbabcd
3
`
	expect := "abacaba+babab+bcbabcb"
	runSample(t, s, expect)
}
