package p5550

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, k int, expect []int) {
	res := decrypt(arr, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{5, 7, 1, 4}
	k := 3
	expect := []int{12, 10, 16, 13}
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 4, 9, 3}
	k := -2
	expect := []int{12, 5, 6, 13}
	runSample(t, arr, k, expect)
}
