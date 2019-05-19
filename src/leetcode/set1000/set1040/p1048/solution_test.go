package p1048

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := longestStrChain(words)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []string{"a", "b", "ba", "bca", "bda", "bdca"}, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, []string{"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh", "zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq", "gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx", "gru"}, 7)
}
