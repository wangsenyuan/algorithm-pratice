package p678

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := checkValidString(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(())((())()()(*)(*()(())())())()()((()())((()))(*", false)
}

func TestSample2(t *testing.T) {
	runSample(t, "()", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "(*))", true)
}
