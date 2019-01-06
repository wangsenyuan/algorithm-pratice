package p970

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, x int, y int, bound int, expect []int) {
	res := powerfulIntegers(x, y, bound)
	sort.Ints(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d, expect %v, but got %v", x, y, bound, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10})
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 5, 15, []int{2, 4, 6, 8, 10, 14})
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 2, 100, []int{2, 3, 5, 9, 17, 33, 65})
}
