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

		if x != (y == "YES") {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 2
1 2
1 3
2 3`
	expect := `NO
NO
YES`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
2 1 2 1 1
1 2
1 3
4 5`
	expect := `NO
NO
YES`
	runSample(t, s, expect)
}
