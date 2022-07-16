package main

import "testing"

func runSample(t *testing.T, A []int, expect string) {

}

func TestSample1(t *testing.T) {
	A := []int{5, 10}
	expect := "CHEF"
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 3}
	expect := "CHEFINA"
	runSample(t, A, expect)
}
