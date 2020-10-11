package p1605

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := countSubgraphsForEachDiameter(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{{1, 2}, {2, 3}, {2, 4}}
	expect := []int{3, 4, 0}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{{1, 2}}
	expect := []int{1}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{{1, 2}, {2, 3}}
	expect := []int{2, 1}
	runSample(t, n, edges, expect)
}
func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{{1, 3}, {1, 4}, {2, 3}}
	expect := []int{3, 2, 1}
	runSample(t, n, edges, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	edges := [][]int{{1, 5}, {2, 3}, {2, 4}, {2, 5}}
	expect := []int{4, 5, 3, 0}
	runSample(t, n, edges, expect)
}
