package p2798

import "testing"

func runSample(t *testing.T, low string, hi string, expect int) {
	res := countSteppingNumbers(low, hi)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	low := "1"
	high := "11"
	expect := 10
	runSample(t, low, high, expect)
}

func TestSample2(t *testing.T) {
	low := "90"
	high := "101"
	expect := 2
	runSample(t, low, high, expect)
}

func TestSample3(t *testing.T) {
	low := "1"
	high := "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	expect := 125046265
	runSample(t, low, high, expect)
}
