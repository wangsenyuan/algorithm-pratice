package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, n, _, d := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var cur int
	for cur+d <= n {
		i := cur
		cur += d
		for cur > i && res[cur-1] == 0 {
			cur--
		}
		if cur == i {
			t.Fatalf("Sample result %v, not correct, not able to jump at %d", res, i+d)
		}
		x := res[cur-1]
		for cur+1 <= n && res[cur] == x {
			cur++
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 3 2
1 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 1 11
1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 1 5
2
`
	expect := true
	runSample(t, s, expect)
}
