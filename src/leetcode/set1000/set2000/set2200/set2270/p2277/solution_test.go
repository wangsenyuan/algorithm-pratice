package p2277

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	sys := Constructor(2, 5)
	res := sys.Gather(4, 0)

	expect(t, []int{0, 0}, res, "result expect %v, but got %v")

	res = sys.Gather(2, 0)

	expect(t, nil, res, "result expect %v, but got %v")

	res1 := sys.Scatter(5, 1)
	if !res1 {
		t.Error("Should have enough seats to scatter 5 persons")
	}

	res2 := sys.Scatter(5, 1)

	if res2 {
		t.Error("Should not have enough seats")
	}
}

func expect(t *testing.T, exp []int, res []int, msg string) {
	if !reflect.DeepEqual(res, exp) {
		t.Errorf(msg, exp, res)
	}
}

func TestSample2(t *testing.T) {
	sys := Constructor(5, 9)
	res1 := sys.Scatter(2, 2)
	// first row taken up by 2
	if !res1 {
		t.Fatal("should take first row, two seats")
	}
	res2 := sys.Gather(7, 2)
	if !reflect.DeepEqual(res2, []int{0, 2}) {
		t.Fatal("should take first row, 7 seats from 2")
	}
}
