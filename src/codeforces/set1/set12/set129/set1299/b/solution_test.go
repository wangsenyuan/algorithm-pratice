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
	s := `4
1 0
4 1
3 4
0 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
0 0
1 0
2 1
3 3
4 6
3 6
2 5
1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
156187286 -987727459
913926112 405880599
884444542 466645316
-345070828 938576647
`
	expect := false
	runSample(t, s, expect)
}
