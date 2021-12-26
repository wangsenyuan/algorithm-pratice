package p2125

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, expect []int64) {
	res := getDistances(arr)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 1, 3, 1, 2, 3, 3}
	expect := []int64{4, 2, 7, 2, 4, 4, 5}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{10, 5, 10, 10}
	expect := []int64{5, 0, 3, 4}
	runSample(t, arr, expect)
}
