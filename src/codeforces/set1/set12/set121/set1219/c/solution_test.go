package main

import "testing"

func runSample(t *testing.T, L int, A string, expect string) {
	res := solve(L, A)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := 3
	A := "123456"
	expect := "124124"
	runSample(t, L, A, expect)
}
func TestSample2(t *testing.T) {
	L := 3
	A := "12345"
	expect := "100100"
	runSample(t, L, A, expect)
}

func TestSample3(t *testing.T) {
	L := 3
	A := "999999"
	expect := "100100100"
	runSample(t, L, A, expect)
}
