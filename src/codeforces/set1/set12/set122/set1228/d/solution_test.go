package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	// already verified in the solution
}

func TestSample1(t *testing.T) {
	s := `6 11
1 2
1 3
1 4
1 5
1 6
2 4
2 5
2 6
3 4
3 5
3 6
`

	expect := true

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6
1 2
1 3
1 4
2 3
2 4
3 4
`

	expect := false

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 9
1 4
1 5
1 6
2 4
2 5
2 6
3 4
3 5
3 6
`

	expect := false

	runSample(t, s, expect)
}
