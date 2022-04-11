package p2230

import "testing"

func runSample(t *testing.T, num int, expect int) {
	res := largestInteger(num)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 1234
	expect := 3412
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 65875
	expect := 87655
	runSample(t, num, expect)
}
