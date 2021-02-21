package p1768

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, box string, expect []int) {
	res := minOperations(box)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", box, expect, res)
	}
}

func TestSample1(t *testing.T) {
	box := "110"
	expect := []int{1, 1, 3}
	runSample(t, box, expect)
}

func TestSample2(t *testing.T) {
	box := "001011"
	expect := []int{11, 8, 5, 4, 3, 4}
	runSample(t, box, expect)
}
