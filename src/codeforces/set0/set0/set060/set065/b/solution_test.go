package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	if !sort.StringsAreSorted(res) {
		t.Fatalf("Sample result %v, not correct", res)
	}
	n := len(res)
	if res[0] < "1000" || res[n-1] > "2011" {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1875
1936
1721
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
9999
2000
3000
3011
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1999
5055
2000
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
2037
2025
`
	expect := true
	runSample(t, s, expect)
}
