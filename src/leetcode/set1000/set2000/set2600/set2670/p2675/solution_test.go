package p2675

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := circularGameLosers(n, k)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	expect := []int{4, 5}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 4
	expect := []int{2, 3, 4}
	runSample(t, n, k, expect)
}
