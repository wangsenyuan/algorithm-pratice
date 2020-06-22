package p1488

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rains []int, expect []int) {
	res := avoidFlood(rains)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", rains, expect, res)
	}
}

func TestSample1(t *testing.T) {
	rains := []int{1, 2, 3, 4}
	expect := []int{-1, -1, -1, -1}
	runSample(t, rains, expect)
}

func TestSample2(t *testing.T) {
	rains := []int{1, 2, 0, 0, 2, 1}
	expect := []int{-1, -1, 2, 1, -1, -1}
	runSample(t, rains, expect)
}

func TestSample3(t *testing.T) {
	rains := []int{1, 2, 0, 1, 2}
	runSample(t, rains, nil)
}

func TestSample4(t *testing.T) {
	rains := []int{69, 0, 0, 0, 69}
	expect := []int{-1, 69, -1, -1, -1}
	runSample(t, rains, expect)
}
