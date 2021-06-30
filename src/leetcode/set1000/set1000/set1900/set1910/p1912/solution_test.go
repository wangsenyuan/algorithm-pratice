package p1912

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	n := 3
	entries := [][]int{
		{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5},
	}

	sys := Constructor(n, entries)

	res := sys.Search(1)

	expect := []int{1, 0, 2}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v, when search 1", expect, res)
	}

	sys.Rent(0, 1)

	sys.Rent(1, 2)

	rep := sys.Report()
	exp := [][]int{
		{0, 1}, {1, 2},
	}

	if !reflect.DeepEqual(rep, exp) {
		t.Fatalf("Sample expect %v, but got %v, when report", exp, rep)
	}

	sys.Drop(1, 2)

	res = sys.Search(2)

	expect = []int{0, 1}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v, when search 2", expect, res)
	}
}
