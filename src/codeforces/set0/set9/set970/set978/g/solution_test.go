package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, n, exams := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	m := len(exams)
	prepare := make([]int, m)

	publishAt := make([][]int, n+1)
	eventAt := make([][]int, n+1)
	for i, exam := range exams {
		s, d, c := exam[0], exam[1], exam[2]
		publishAt[s] = append(publishAt[s], i)
		eventAt[d] = append(eventAt[d], i)
		prepare[i] = c
	}

	for i := 1; i <= n; i++ {
		if len(eventAt[i]) == 1 {
			j := eventAt[i][0]
			if prepare[j] > 0 || res[i-1] != m+1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}

		for _, j := range publishAt[i] {
			c := exams[j][2]
			if prepare[j] != c {
				t.Fatalf("Sample result %v, not correct, starting preparing before publishing", res)
			}
		}
		if res[i-1] > 0 && res[i-1] <= m {
			prepare[res[i-1]-1]--
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
1 3 1
1 5 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 3 1
1 2 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3
4 7 2
1 10 3
8 9 1
`
	expect := true
	runSample(t, s, expect)
}
