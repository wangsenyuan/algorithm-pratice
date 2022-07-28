package main

import "testing"

func runSample(t *testing.T, F []int, C []int, expect int) {

}

func TestSample1(t *testing.T) {
	F := []int{1, 3}
	C := []int{2, 4}
	expect := 3
	runSample(t, F, C, expect)
}

func TestSample2(t *testing.T) {
	F := []int{100, 200, 300}
	C := []int{1}
	expect := 2
	runSample(t, F, C, expect)
}

func TestSample3(t *testing.T) {
	F := []int{1}
	C := []int{100, 200}
	expect := 1
	runSample(t, F, C, expect)
}
