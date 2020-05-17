package p1450

import "testing"

func runSample(t *testing.T, text string, expect string) {
	res := arrangeWords(text)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", text, expect, res)
	}
}

func TestSample1(t *testing.T) {
	text := "Leetcode is cool"
	expect := "Is cool leetcode"
	runSample(t, text, expect)
}
