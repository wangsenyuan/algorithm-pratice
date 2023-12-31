package p2982

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []bool) {
	res := canMakePalindromeQueries(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcabc"
	queries := [][]int{
		{1, 1, 3, 5}, {0, 2, 5, 5},
	}
	expect := []bool{true, true}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "abbcdecbba"
	queries := [][]int{
		{0, 2, 7, 9},
	}
	expect := []bool{false}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "acbcab"
	queries := [][]int{
		{1, 2, 4, 5},
	}
	expect := []bool{true}
	runSample(t, s, queries, expect)
}

func TestSample4(t *testing.T) {
	s := "dcdadadc"
	queries := [][]int{
		{2, 3, 6, 7},
	}
	expect := []bool{true}
	runSample(t, s, queries, expect)
}
