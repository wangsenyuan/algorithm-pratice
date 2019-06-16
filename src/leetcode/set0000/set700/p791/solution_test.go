package p791

import "testing"

func runSample(t *testing.T, S, T string, expect string) {
	res := customSortString(S, T)

	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", S, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "cba", "abcd", "cbad")
}
