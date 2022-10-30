package p2457

import "testing"

func runSample(t *testing.T, num int64, target int, expect int64) {
	res := makeIntegerBeautiful(num, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var num int64 = 16
	target := 6
	var expect int64 = 4
	runSample(t, num, target, expect)
}

func TestSample2(t *testing.T) {
	var num int64 = 467
	target := 6
	var expect int64 = 33
	runSample(t, num, target, expect)
}
func TestSample3(t *testing.T) {
	var num int64 = 1
	target := 1
	var expect int64 = 0
	runSample(t, num, target, expect)
}
