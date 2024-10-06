package p3311

import (
	"reflect"
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int) {
	res := constructGridLayout(n, edges)

	var resEdges [][]int
	h := len(res)
	w := len(res[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i > 0 {
				resEdges = append(resEdges, []int{res[i][j], res[i-1][j]})
			}
			if j > 0 {
				resEdges = append(resEdges, []int{res[i][j], res[i][j-1]})
			}
		}
	}
	if len(resEdges) != len(edges) {
		t.Fatalf("Sample result %v, got wrong number of edges", res)
	}
	for _, cur := range edges {
		if cur[0] > cur[1] {
			cur[0], cur[1] = cur[1], cur[0]
		}
	}
	for _, cur := range resEdges {
		if cur[0] > cur[1] {
			cur[0], cur[1] = cur[1], cur[0]
		}
	}

	slices.SortFunc(edges, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	slices.SortFunc(resEdges, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	if !reflect.DeepEqual(edges, resEdges) {
		t.Fatalf("Sample result %v, got wrong edges %v, expect %v", res, resEdges, edges)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	runSample(t, n, edges)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1}, {1, 3}, {2, 3}, {2, 4}}
	runSample(t, n, edges)
}

func TestSample3(t *testing.T) {
	n := 9
	edges := [][]int{{0, 1}, {0, 4}, {0, 5}, {1, 7}, {2, 3}, {2, 4}, {2, 5}, {3, 6}, {4, 6}, {4, 7}, {6, 8}, {7, 8}}
	runSample(t, n, edges)
}
