package main

import "testing"

func runSample(t *testing.T, s string, x string, expect string) {
	res := solve(s, x)

	a := countOccurrence(expect, x)
	b := countOccurrence(res, x)
	if a != b {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func countOccurrence(s string, x string) int {
	var res int
	for i := 0; i+len(x) <= len(s); i++ {
		if s[i:i+len(x)] == x {
			res++
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	s := "101101"
	x := "110"
	expect := "110110"
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "10010110"
	x := "100011"
	expect := "01100011"
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "10101010"
	x := "1001"
	expect := "10010011"
	runSample(t, s, x, expect)
}
