package main

import "testing"

func runSample(t *testing.T, A int64, S int64, expect int64) {
	res := solve(A, S)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var a int64 = 17236
	var s int64 = 1106911
	var expect int64 = 3465
	runSample(t, a, s, expect)
}

func TestSample2(t *testing.T) {
	var a int64 = 1
	var s int64 = 5
	var expect int64 = 4
	runSample(t, a, s, expect)
}

func TestSample3(t *testing.T) {
	var a int64 = 108
	var s int64 = 112
	var expect int64 = -1
	runSample(t, a, s, expect)
}
func TestSample4(t *testing.T) {
	var a int64 = 12345
	var s int64 = 1023412
	var expect int64 = 90007
	runSample(t, a, s, expect)
}

func TestSample5(t *testing.T) {
	var a int64 = 1
	var s int64 = 11
	var expect int64 = 10
	runSample(t, a, s, expect)
}

func TestSample6(t *testing.T) {
	var a int64 = 1
	var s int64 = 20
	var expect int64 = -1
	runSample(t, a, s, expect)
}
