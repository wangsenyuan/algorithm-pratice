package p3327

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, parent []int, s string, expect []bool) {
	res := findAnswer(parent, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parent := []int{-1, 0, 0, 1, 1, 2}
	s := "aababa"
	expect := []bool{true, true, false, true, true, true}
	runSample(t, parent, s, expect)
}

func TestSample2(t *testing.T) {
	parent := []int{-1, 0, 0, 0, 0}
	s := "aabcb"
	expect := []bool{true, true, true, true, true}
	runSample(t, parent, s, expect)
}

func TestSample3(t *testing.T) {
	parent := []int{-1, 4, 4, 1, 0, 2, 4}
	s := "bfgiigf"
	expect := []bool{false, false, true, true, true, true, true}
	runSample(t, parent, s, expect)
}

func TestSample4(t *testing.T) {
	parent := []int{-1, 6, 13, 1, 1, 3, 0, 0, 14, 3, 14, 13, 16, 3, 9, 6, 9, 14, 17}
	s := "ehdfjgieggkddkdfkdd"
	expect := []bool{false, false, true, false, true, true, false, true, true, true, true, true, true, false, false, true, false, true, true}
	runSample(t, parent, s, expect)
}

func TestSample5(t *testing.T) {
	parent := []int{-1, 0}
	s := "aa"
	expect := []bool{true, true}
	runSample(t, parent, s, expect)
}
