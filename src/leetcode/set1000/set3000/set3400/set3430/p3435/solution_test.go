package p3435

import (
	"reflect"
	"slices"
	"testing"
)

func runSample(t *testing.T, words []string, expect [][]int) {
	res := supersequences(words)

	cmp := func(a, b []int) int {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return -1
			}
			if a[i] > b[i] {
				return 1
			}
		}
		return 0
	}

	slices.SortFunc(expect, cmp)

	slices.SortFunc(res, cmp)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"ab", "ba"}
	expect := [][]int{
		{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"aa", "ac"}
	expect := [][]int{
		{2, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"aa", "bb", "cc"}
	expect := [][]int{
		{2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{"jq", "iq", "qj", "qi", "ij"}
	expect := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	runSample(t, words, expect)
}
