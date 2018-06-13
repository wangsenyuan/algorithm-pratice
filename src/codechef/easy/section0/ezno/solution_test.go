package main

import "testing"

func TestSample1(t *testing.T) {
	res := solve(2, 2)
	if res != 1 {
		t.Errorf("sample 2 2, expect 1, but got %d", res)
	}
}
