package p819

import "testing"

func runSample(t *testing.T, p string, ban []string, expected string) {
	res := mostCommonWord(p, ban)

	if res != expected {
		t.Errorf("sample %s %v, expect %s, but got %s", p, ban, expected, res)
	}
}

func TestSample1(t *testing.T) {
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	ban := []string{"hit"}
	expected := "ball"
	runSample(t, p, ban, expected)
}
