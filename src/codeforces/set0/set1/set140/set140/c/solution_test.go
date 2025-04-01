package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, r := process(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	cnt := make(map[int]int)
	for _, v := range r {
		cnt[v]++
	}
	for _, cur := range res {
		if cur[0] <= cur[1] || cur[1] <= cur[2] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		for _, v := range cur {
			if cnt[v] == 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
			cnt[v]--
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7
1 2 3 4 5 6 7
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 2 3
`
	expect := 0
	runSample(t, s, expect)
}
