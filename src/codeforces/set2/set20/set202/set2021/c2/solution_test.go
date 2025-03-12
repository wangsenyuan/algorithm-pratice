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
		if y == "YA" != x {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 2 2
1 2 3 4
1 1
1 2
1 1`
	expect := `YA
TIDAK
YA`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6 2
1 2 3
1 1 2 3 3 2
3 3
2 2`
	expect := `YA
TIDAK
YA`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 6 2
3 1 4 2
3 1 1 2 3 4
3 4
4 2`
	expect := `TIDAK
YA
YA`
	runSample(t, s, expect)
}
