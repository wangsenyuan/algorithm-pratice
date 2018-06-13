package main

import "testing"

func TestSampleWithSolution1(t *testing.T) {
	hs := []int{1, 2, 3, 2, 3}
	res := solve1(5, hs)
	if res != 3 {
		t.Errorf("there should be 3 points cross %v, but get %d", hs, res)
	}
}

func TestSample1WithSolution1(t *testing.T) {
	hs := []int{1, 1, 1}
	res := solve1(1, hs)
	if res != 1 {
		t.Errorf("there should be 3 points cross %v, but get %d", hs, res)
	}
}

func TestSample(t *testing.T) {
	hs := []int{1, 2, 3, 2, 3}
	res := solve(5, hs)
	if res != 3 {
		t.Errorf("there should be 3 points cross %v, but get %d", hs, res)
	}
}

func TestSample1WithSolution(t *testing.T) {
	hs := []int{1, 1, 1}
	res := solve(1, hs)
	if res != 1 {
		t.Errorf("there should be 3 points cross %v, but get %d", hs, res)
	}
}
