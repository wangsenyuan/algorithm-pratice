package main

import "testing"

func runSample(t *testing.T, num string, k int, expect string) {
	res := minInteger(num, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "4321"
	k := 4
	expect := "1342"
	runSample(t, num, k, expect)
}

func TestSample2(t *testing.T) {
	num := "100"
	k := 1
	expect := "010"
	runSample(t, num, k, expect)
}

func TestSample3(t *testing.T) {
	num := "36789"
	k := 1000
	expect := "36789"
	runSample(t, num, k, expect)
}

func TestSample5(t *testing.T) {
	num := "9438957234785635408"
	k := 23
	expect := "0345989723478563548"
	runSample(t, num, k, expect)
}
