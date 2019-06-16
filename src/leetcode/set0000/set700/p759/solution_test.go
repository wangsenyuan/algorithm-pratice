package p759

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	input := [][]Interval{
		[]Interval{{1, 2}, {5, 6}},
		[]Interval{{1, 3}},
		[]Interval{{4, 10}},
	}

	res := employeeFreeTime(input)
	expect := []Interval{{3, 4}}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v, should give %v, but got %v", input, expect, res)
	}
}
