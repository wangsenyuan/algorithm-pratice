package p2827

import "testing"

func runSample(t *testing.T, low int, high int, k int, expect int) {
	res := numberOfBeautifulIntegers(low, high, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	low := 10
	high := 20
	k := 3
	expect := 2
	runSample(t, low, high, k, expect)
}

func TestSample2(t *testing.T) {
	low := 1
	high := 10
	k := 1
	expect := 1
	runSample(t, low, high, k, expect)
}

func TestSample3(t *testing.T) {
	low := 5
	high := 5
	k := 2
	expect := 0
	runSample(t, low, high, k, expect)
}
