package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	a, res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(res)
	for i := 0; i < n; i++ {
		if res[i] == res[(i+1)%n] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		a[res[i]-1]--
		if a[res[i]-1] < 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
1 3 5
`
	expect := true
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `10 2
5 5
`
	expect := true
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `10 3
1 10 3
`
	expect := false
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `6 3
2 2 2
`
	expect := true
	runSample(t, s, expect)
}
