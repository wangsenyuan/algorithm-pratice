package main

import "testing"

func runUniqueSubStrCount(t *testing.T, s string, expect int64) {
	res := uniqueSubStrCount([]byte(s))
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestUniqueSubStrCount1(t *testing.T) {
	runUniqueSubStrCount(t, "abaabba", 21)
}

func runSample(t *testing.T, a, b string, expect int64) {
	res := solve([]byte(a), []byte(b))

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aacd", "cdaa", 8)
}
