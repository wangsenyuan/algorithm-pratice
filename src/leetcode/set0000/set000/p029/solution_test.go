package p029

import "testing"

func TestSample1(t *testing.T) {
	res := divide(-1, 1)
	if res != -1 {
		t.Errorf("sample -1 1 should give -1, but got %d", res)
	}
}
