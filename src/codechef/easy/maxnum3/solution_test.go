package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s []int, expect []int) {
	can, res := solve(s)
	if !can && len(expect) != 0 {
		t.Errorf("sample %v, expect %v, but got wrong answer", s, expect)
	} else if can && !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []int{1, 2, 3}
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []int{1, 0, 0, 2, 2, 2}
	expect := []int{0, 0, 2, 2, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []int{1, 0, 0, 0, 3}
	expect := []int{}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []int{1, 0, 2, 2}
	expect := []int{1, 0, 2}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := []int{1, 3, 2, 3}
	expect := []int{1, 3, 2}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := []int{1, 2, 3, 2, 3, 6}
	expect := []int{1, 3, 2, 3, 6}
	runSample(t, s, expect)
}
