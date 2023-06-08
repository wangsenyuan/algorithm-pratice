package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	var open int
	for i := 0; i < len(s); i++ {
		if res[i] == '(' {
			open++
		} else {
			open--
		}
		if open < 0 {
			t.Fatalf("Sample result %s, not correct", res)
		}
		if i > 0 {
			if s[i] == s[i-1] && res[i] != res[i-1] {
				t.Fatalf("Sample result %s, not correct", res)
			}
		}
	}

	if open != 0 {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abbb"
	expect := false
	runSample(t, s, expect)
}
