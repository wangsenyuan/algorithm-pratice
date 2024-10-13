package p3321

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, k int, x int, expect []int64) {
	res := findXSum(a, k, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 2, 2, 3, 4, 2, 3}
	k := 6
	x := 2
	expect := []int64{6, 10, 12}
	runSample(t, a, k, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 8, 7, 8, 7, 5}
	k := 2
	x := 2
	expect := []int64{11, 15, 15, 15, 12}
	runSample(t, a, k, x, expect)
}
