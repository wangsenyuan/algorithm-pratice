package p1362

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, num int, expect []int) {
	res := closestDivisors(num)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d, expect %v, but got %v", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, []int{3, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, 123, []int{5, 25})
}

func TestSample3(t *testing.T) {
	runSample(t, 999, []int{25, 40})
}
