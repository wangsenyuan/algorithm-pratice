package p745

import "testing"

func TestSample1(t *testing.T) {
	filter := Constructor([]string{"apple"})
	res := filter.F("a", "e")
	if res != 0 {
		t.Errorf("find a & e should give 0, but got %d", res)
	}
	res = filter.F("b", "")
	if res != -1 {
		t.Errorf("find a & '' should give -1, but got %d", res)
	}
	res = filter.F("a", "le")
	if res != 0 {
		t.Errorf("find a & le should give 0, but got %d", res)
	}

	res = filter.F("apple", "apple")
	if res != 0 {
		t.Errorf("find apple & apple should give 0, but got %d", res)
	}
}
