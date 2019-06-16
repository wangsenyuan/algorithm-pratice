package p754

import "testing"

func TestSample1(t *testing.T) {
	target := 3
	res := reachNumber(target)
	if res != 2 {
		t.Errorf("sample %d; should give answer 2, but got %d", target, res)
	}
}

func TestSample2(t *testing.T) {
	target := 2
	res := reachNumber(target)
	if res != 3 {
		t.Errorf("sample %d; should give answer 3, but got %d", target, res)
	}
}
