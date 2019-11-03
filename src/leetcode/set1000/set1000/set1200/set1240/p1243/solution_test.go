package p1243

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, expect []int) {
	res := transformArray(arr)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{6, 2, 3, 4}
	expect := []int{6, 3, 3, 4}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 6, 3, 4, 3, 5}
	expect := []int{1, 4, 4, 4, 4, 5}
	runSample(t, arr, expect)
}
