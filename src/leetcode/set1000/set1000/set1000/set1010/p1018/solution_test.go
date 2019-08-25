package p1018

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []bool) {
	res := prefixesDivBy5(A)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{0, 1, 1}, []bool{true, false, false})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 1}, []bool{false, false, false})
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 1, 1, 1, 1, 1}, []bool{true, false, false, false, true, false})
}

func TestSample4(t *testing.T) {
	runSample(t, []int{1, 1, 1, 0, 1}, []bool{false, false, false, false, false})
}
