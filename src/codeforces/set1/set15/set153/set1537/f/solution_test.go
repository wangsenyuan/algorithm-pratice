package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
5 1 2 -3
3 3 10 1
1 2
1 4
3 2
3 4`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
5 8 6 6
-3 1 15 4
1 2
1 4
3 2
3 4`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 10
521 696 891 558 787
-902 275 -957 818 531
2 3
5 4
1 4
1 3
2 4
1 2
5 3
2 5
1 5
4 3`
	expect := true
	runSample(t, s, expect)
}
