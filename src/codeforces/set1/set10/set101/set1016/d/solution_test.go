package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, a, b := process(reader)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	n := len(res)
	m := len(res[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[i] ^= res[i][j]
			b[j] ^= res[i][j]
		}
	}
	if slices.Max(a) > 0 || slices.Max(b) > 0 {
		t.Errorf("Sample result is not correct")
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
2 9
5 3 13
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 7 6
2 15 12
`
	expect := false
	runSample(t, s, expect)
}
