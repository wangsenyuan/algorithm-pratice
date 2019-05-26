package p1053

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := prevPermOpt1(A)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{3, 2, 1}, []int{3, 1, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 5}, []int{1, 1, 5})
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1, 9, 4, 6, 7}, []int{1, 7, 4, 6, 9})
}

func TestSample4(t *testing.T) {
	runSample(t, []int{3, 1, 1, 3}, []int{1, 3, 1, 3})
}
