package p1719

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, encoded []int, first int, expect []int) {
	res := decode(encoded, first)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	encoded := []int{1, 2, 3}
	first := 1
	expect := []int{1, 0, 2, 1}
	runSample(t, encoded, first, expect)
}

func TestSample2(t *testing.T) {
	encoded := []int{6, 2, 7, 3}
	first := 4
	expect := []int{4, 2, 0, 7, 4}
	runSample(t, encoded, first, expect)
}
