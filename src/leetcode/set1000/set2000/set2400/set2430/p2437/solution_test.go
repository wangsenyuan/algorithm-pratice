package p2437

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, queries [][]int, expect []int) {
	res := productQueries(n, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	n := 15
	queries := [][]int{{0, 1}, {2, 2}, {0, 3}}
	expect := []int{2, 4, 64}
	runSample(t, n, queries, expect)
}
