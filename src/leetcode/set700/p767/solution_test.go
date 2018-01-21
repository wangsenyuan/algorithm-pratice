package p767

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := reorganizeString(s)
	if res != expect {
		t.Errorf("sample %s, expect {%s}, but got {%s}", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aab", "aba")
}

func TestSample2(t *testing.T) {
	runSample(t, "aaab", "")
}
