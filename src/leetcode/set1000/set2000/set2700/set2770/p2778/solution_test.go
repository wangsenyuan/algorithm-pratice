package p2778

import "testing"

func runSample(t *testing.T, s string, forbidden []string, expect int) {
	res := longestValidSubstring(s, forbidden)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cbaaaabc"
	words := []string{"aaa", "cb"}
	expect := 4
	runSample(t, s, words, expect)
}

func TestSample2(t *testing.T) {
	s := "leetcode"
	words := []string{"de", "le", "e"}
	expect := 4
	runSample(t, s, words, expect)
}

func TestSample3(t *testing.T) {
	s := "bbccc"
	words := []string{"ccc", "bcba", "bcc"}
	expect := 3
	runSample(t, s, words, expect)
}

func TestSample4(t *testing.T) {
	s := "bcbab"
	words := []string{"cba", "aaab", "bcbc", "caba"}
	expect := 3
	runSample(t, s, words, expect)
}
