package p1977

import "testing"

func runSample(t *testing.T, num string, expect int) {
	res := numberOfCombinations(num)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "327"
	expect := 2
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := "094"
	expect := 0
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	num := "9999999999999"
	expect := 101
	runSample(t, num, expect)
}
