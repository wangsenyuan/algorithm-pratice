package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	reader := bufio.NewReader(strings.NewReader(expect))

	for _, x := range res {
		y := readString(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
b
aa
hg
c
.
`
	expect := `aa.b
c.hg
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
aa
a
!
`
	expect := `a!aa`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
aa
a
|
`
	expect := `aa|a`
	runSample(t, s, expect)
}
