package p2836

import "testing"

func runSample(t *testing.T, receiver []int, k int64, expect int64) {
	res := getMaxFunctionValue(receiver, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	receiver := []int{2, 0, 1}
	var k int64 = 4
	var expect int64 = 6
	runSample(t, receiver, k, expect)
}

func TestSample2(t *testing.T) {
	receiver := []int{1, 1, 1, 2, 3}
	var k int64 = 3
	var expect int64 = 10
	runSample(t, receiver, k, expect)
}

func TestSample3(t *testing.T) {
	receiver := []int{1, 1}
	var k int64 = 2
	var expect int64 = 3
	runSample(t, receiver, k, expect)
}
