package p953

import "testing"

func runSample(t *testing.T, words []string, order string, expect bool) {
	res := isAlienSorted(words, order)
	if res != expect {
		t.Errorf("Sample %v %s, expect %t, but got %t", words, order, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	runSample(t, words, order, true)
}

func TestSample2(t *testing.T) {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"
	runSample(t, words, order, false)
}

func TestSample3(t *testing.T) {
	words := []string{"apple", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"
	runSample(t, words, order, false)
}
