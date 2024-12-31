package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	check := func(s string, res string) int {
		var r []byte
		var b []byte
		for i := range len(s) {
			if res[i] == '0' {
				r = append(r, s[i])
			} else {
				b = append(b, s[i])
			}
		}
		return max(getDepth(string(r)), getDepth(string(b)))
	}

	x := check(s, expect)
	y := check(s, res)

	if x != y {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "()"
	expect := "11"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "(())"
	expect := "0101"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "((()())())"
	expect := "0110001111"
	runSample(t, s, expect)
}
