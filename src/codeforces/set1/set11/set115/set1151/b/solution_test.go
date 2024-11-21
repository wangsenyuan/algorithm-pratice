package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, a := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var sum int

	for i := 0; i < len(res); i++ {
		sum ^= a[i][res[i]-1]
	}

	if sum == 0 {
		t.Fatalf("Sample result %v, leading to xor 0", res)
	}

}

func TestSample1(t *testing.T) {
	s := `3 2
0 0
0 0
0 0
`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
7 7 7
7 7 10
`
	expect := true
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `1 1
7
`
	expect := true
	runSample(t, s, expect)
}
