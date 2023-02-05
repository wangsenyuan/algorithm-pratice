package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{235499701, 451218171, 355604420, 132973458}
	B := []int{365049318, 264083156, 491406845, 62875547, 175951751}
	expect := 531556171
	runSample(t, A, B, expect)
}
