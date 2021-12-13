package p2066

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, restrictions [][]int, requests [][]int, expect []bool) {
	res := friendRequests(n, restrictions, requests)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	rest := [][]int{{0, 1}}
	req := [][]int{{0, 2}, {2, 1}}
	expect := []bool{true, false}
	runSample(t, n, rest, req, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	rest := [][]int{{0, 1}}
	req := [][]int{{1, 2}, {2, 0}}
	expect := []bool{true, false}
	runSample(t, n, rest, req, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	rest := [][]int{{0, 1}, {1, 2}, {2, 3}}
	req := [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}
	expect := []bool{true, false, true, false}
	runSample(t, n, rest, req, expect)
}
