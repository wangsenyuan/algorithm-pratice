package p1040

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, stones []int, expect []int) {
	res := numMovesStonesII(stones)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{7, 4, 9}, []int{1, 2})
}
