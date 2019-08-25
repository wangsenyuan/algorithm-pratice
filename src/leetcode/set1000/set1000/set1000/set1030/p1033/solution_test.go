package p1033

import "testing"

func TestSample1(t *testing.T) {
	res := numMovesStones(4, 3, 2)
	a, b := res[0], res[1]
	if a != 0 || b != 0 {
		t.Errorf("Sample 4, 3, 2, expect (0, 0), but got (%d %d)", a, b)
	}
}

func TestSample2(t *testing.T) {
	res := numMovesStones(3, 5, 1)
	a, b := res[0], res[1]
	if a != 1 || b != 2 {
		t.Errorf("Sample 3, 5, 1, expect (1, 2), but got (%d %d)", a, b)
	}
}
