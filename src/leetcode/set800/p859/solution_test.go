package p859

import "testing"

func runSample(t *testing.T, A, B string, expect bool) {
	res := buddyStrings(A, B)

	if res != expect {
		t.Errorf("sample %s %s, expect %t, but got %t", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ab", "ba", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "ab", "ab", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "aa", "aa", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "aaaaaaabc", "aaaaaaacb", true)
}

func TestSample5(t *testing.T) {
	runSample(t, "", "", false)
}
