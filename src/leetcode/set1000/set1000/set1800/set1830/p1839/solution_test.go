package p1839

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := longestBeautifulSubstring(word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aeiaaioaaaaeiiiiouuuooaauuaeiu"
	expect := 13
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aeeeiiiioooauuuaeiou"
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "a"
	expect := 0
	runSample(t, s, expect)
}
