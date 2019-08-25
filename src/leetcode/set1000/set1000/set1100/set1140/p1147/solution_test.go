package p1147

import "testing"

func runSample(t *testing.T, text string, expect int) {
	res := longestDecomposition(text)
	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", text, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ghiabcdefhelloadamhelloabcdefghi", 7)
}

func TestSample2(t *testing.T) {
	runSample(t, "merchant", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "antaprezatepzapreanta", 11)
}

func TestSample4(t *testing.T) {
	runSample(t, "aaa", 3)
}
