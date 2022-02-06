package p2160

import (
	"testing"
)

func runSample(t *testing.T, num int, expect int) {
	res := minimumSum(num)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2932, 52)
}


func TestSample2(t *testing.T) {
	runSample(t, 4009, 13)
}