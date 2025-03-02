package p3470

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := permute(n, int64(k))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 6
	expect := []int{3, 4, 1, 2}
	runSample(t, n, k, expect)
}
