package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	r := bufio.NewReader(strings.NewReader(s))

	res, p, q, k := process(r)

	if len(res) > 0 != expect {
		t.Fatalf("Sample %s, expect %t, but got %s", s, expect, res)
	}

	if !expect {
		return
	}
	freq := make(map[byte]int)
	for i := 0; i < len(res); i++ {
		freq[res[i]]++
	}
	if len(freq) < k {
		t.Fatalf("Sample result %s, not having %d distinct letters", s, k)
	}
	for i := 0; i+1 < len(res); i++ {
		u, v := p[i]-1, p[i+1]-1
		if res[u] > res[v] {
			t.Fatalf("Sample result %s not correct, it violates condition at %d", res, i)
		}
		u, v = q[i]-1, q[i+1]-1
		if res[u] > res[v] {
			t.Fatalf("Sample result %s not correct, it violates condition at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2 3
1 3 2
`
	expect := true
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `2 1
1 2
2 1
`
	expect := true
	runSample(t, s, expect)
}
