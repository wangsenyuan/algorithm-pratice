package p761

import "testing"

func TestSample1(t *testing.T) {
	S := "11011000"
	E := "11100100"
	R := makeLargestSpecial(S)
	if E != R {
		t.Errorf("sample %s, should give %s, but got %s", S, E, R)
	}
}
