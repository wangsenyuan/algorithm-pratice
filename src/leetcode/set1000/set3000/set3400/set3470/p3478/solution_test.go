package p3478

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, b []int, k int, expect []int64) {
	res := findMaxSum(a, b, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 1, 5, 3}
	b := []int{10, 20, 30, 40, 50}
	k := 2
	expect := []int64{80, 30, 0, 80, 50}
	runSample(t, a, b, k, expect)
}
