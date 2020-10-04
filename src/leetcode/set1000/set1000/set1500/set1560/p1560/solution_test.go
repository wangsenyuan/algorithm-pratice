package p1560

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, rounds []int, expect []int) {
	res := mostVisited(n, rounds)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	rounds := []int{1, 3, 1, 2}
	expect := []int{1, 2}
	runSample(t, n, rounds, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	rounds := []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	expect := []int{2}
	runSample(t, n, rounds, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	rounds := []int{1, 3, 5, 7}
	expect := []int{1, 2, 3, 4, 5, 6, 7}
	runSample(t, n, rounds, expect)
}
