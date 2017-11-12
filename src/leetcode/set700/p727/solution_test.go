package p727

import "testing"

func TestSample1(t *testing.T) {
	S := "abcdebdde"
	T := "bde"
	res := minWindow(S, T)
	exp := "bcde"
	if res != exp {
		t.Errorf("minWindow(%s, %s) should be %s, but got %s", S, T, exp, res)
	}
}
