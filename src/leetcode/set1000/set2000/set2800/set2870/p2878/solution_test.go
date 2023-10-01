package p2878

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges []int, expect []int) {
	res := countVisitedNodes(edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := []int{1, 2, 0, 0}
	expect := []int{3, 3, 3, 4}
	runSample(t, edges, expect)
}

func TestSample2(t *testing.T) {
	edges := []int{1, 2, 3, 4, 0}
	expect := []int{5, 5, 5, 5, 5}
	runSample(t, edges, expect)
}

func TestSample3(t *testing.T) {
	edges := []int{6, 3, 6, 1, 0, 8, 0, 6, 6}
	expect := []int{2, 2, 3, 2, 3, 4, 2, 3, 3}
	runSample(t, edges, expect)
}

func TestSample4(t *testing.T) {
	edges := []int{3, 6, 1, 0, 5, 7, 4, 3}
	expect := []int{2, 7, 8, 2, 5, 4, 6, 3}
	runSample(t, edges, expect)
}
