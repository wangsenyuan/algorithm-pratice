package main

import "testing"

func runSample(t *testing.T, n int, s string, x string, expect bool) {
	res := solve(n, s, x)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	if len(res) != 3*n {
		t.Fatalf("Sample result %s, not correct", res)
	}

	for i := 0; i < 3*n; i++ {
		if res[i] != 'a' && res[i] != 'b' && res[i] != 'c' {
			t.Fatalf("Sample result %s, not correct, it contains non-abc", res)
		}
	}

	for i := 0; i+1 < n; i++ {
		if res[i:i+2] == s || res[i:i+2] == x {
			t.Fatalf("Sample result %s, not correct, it contains s/t at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s := "ab"
	x := "bc"
	expect := true
	runSample(t, n, s, x, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	s := "aa"
	x := "bc"
	expect := true
	runSample(t, n, s, x, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	s := "ab"
	x := "ac"
	expect := true
	runSample(t, n, s, x, expect)
}

func TestSample4(t *testing.T) {
	n := 100
	s := "ba"
	x := "ab"
	expect := true
	runSample(t, n, s, x, expect)
}
