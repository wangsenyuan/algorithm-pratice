package p1104

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, label int, expect []int) {
	res := pathInZigZagTree(label)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d, expect %v, but got %v", label, expect, res)
	}
}

func TestSample1(t *testing.T) {
	label := 14
	expect := []int{1, 3, 4, 14}
	runSample(t, label, expect)
}

func TestSample2(t *testing.T) {
	label := 26
	expect := []int{1, 2, 6, 10, 26}
	runSample(t, label, expect)
}
