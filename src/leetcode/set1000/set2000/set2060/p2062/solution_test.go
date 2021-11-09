package p2062

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := countVowelSubstrings(word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "aeiouu"
	expect := 2
	runSample(t, word, expect)
}

func TestSample2(t *testing.T) {
	word := "cuaieuouac"
	expect := 7
	runSample(t, word, expect)
}
