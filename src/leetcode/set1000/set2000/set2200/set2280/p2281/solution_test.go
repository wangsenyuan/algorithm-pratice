package p2281

import "testing"

func runSample(t *testing.T, strength []int, expect int) {
	res := totalStrength(strength)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	strength := []int{1, 3, 1, 2}
	expect := 44
	runSample(t, strength, expect)
}

func TestSample2(t *testing.T) {
	strength := []int{5, 4, 6}
	expect := 213
	runSample(t, strength, expect)
}
