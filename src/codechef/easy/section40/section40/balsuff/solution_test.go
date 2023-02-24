package main

import "testing"

func runSample(t *testing.T, k int, S string, expect string) {
	res := solve(S, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := "aaa"
	k := 1
	expect := "aaa"
	runSample(t, k, s, expect)
}

func TestSample2(t *testing.T) {
	s := "baba"
	k := 2
	expect := "aabb"
	runSample(t, k, s, expect)
}

func TestSample3(t *testing.T) {
	s := "babb"
	k := 1
	expect := ""
	runSample(t, k, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abcbcac"
	k := 2
	expect := "abcabcc"
	runSample(t, k, s, expect)
}
