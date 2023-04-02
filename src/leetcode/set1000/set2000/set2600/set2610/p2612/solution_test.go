package p2612

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, banned []int, p int, k int, expect []int) {
	res := minReverseOperations(n, p, banned, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	p := 0
	banned := []int{1, 2}
	k := 4
	expect := []int{0, -1, -1, 1}
	runSample(t, n, banned, p, k, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	p := 0
	banned := []int{2, 4}
	k := 3
	expect := []int{0, -1, -1, -1, -1}
	runSample(t, n, banned, p, k, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	p := 2
	banned := []int{0, 1, 3}
	k := 1
	expect := []int{-1, -1, 0, -1}
	runSample(t, n, banned, p, k, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	p := 1
	banned := []int{}
	k := 2
	expect := []int{1, 0}
	runSample(t, n, banned, p, k, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	p := 0
	banned := []int{}
	k := 4
	expect := []int{0, -1, -1, 1}
	runSample(t, n, banned, p, k, expect)
}

func TestSample6(t *testing.T) {
	n := 4
	p := 2
	banned := []int{}
	k := 4
	expect := []int{-1, 1, 0, -1}
	runSample(t, n, banned, p, k, expect)
}

func TestSample7(t *testing.T) {
	n := 5
	p := 0
	banned := []int{}
	k := 4
	expect := []int{0, 3, 2, 1, 4}
	runSample(t, n, banned, p, k, expect)
}
