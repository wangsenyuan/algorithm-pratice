package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {

	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
3
1 6
1 7
1 5
2
1 4
1 3
3
1 10
3 5
2 3
3
1 15
2 4
1 10
1
1 100
`

	expect := 263
	runSample(t, s, expect)

}
