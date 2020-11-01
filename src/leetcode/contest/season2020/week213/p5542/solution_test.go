package p5542

import "testing"

func runSample(tc *testing.T, words []string, target string, expect int) {
	res := numWays(words, target)

	if res != expect {
		tc.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"acca", "bbbb", "caca"}
	target := "aba"
	expect := 6
	runSample(t, words, target, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"abba", "baab"}
	target := "bab"
	expect := 4
	runSample(t, words, target, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"abcd"}
	target := "abcd"
	expect := 1
	runSample(t, words, target, expect)
}

func TestSample4(t *testing.T) {
	words := []string{"abab", "baba", "abba", "baab"}
	target := "abba"
	expect := 16
	runSample(t, words, target, expect)
}
