package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	for _, lyric := range res {
		first := lyric[0]
		second := lyric[1]
		s1 := strings.Split(first, " ")
		s2 := strings.Split(second, " ")
		c11 := convert(s1[0])
		c21 := convert(s2[0])
		if c11.cnt != c21.cnt {
			t.Fatalf("Sample result %v, not correct", res)
		}
		c12 := convert(s1[1])
		c22 := convert(s2[1])
		if c12 != c22 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `14
wow
this
is
the
first
mcdics
codeforces
round
hooray
i
am
proud
about
that
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
arsijo
suggested
the
idea
for
this
problem
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
same
same
same
differ
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1
impossible
`
	expect := 0
	runSample(t, s, expect)
}
