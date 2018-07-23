package p828

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := uniqueLetterString(S)
	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ABC", 10)
}

func TestSample2(t *testing.T) {
	runSample(t, "ABA", 8)
}
