package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "IIpqDpNp", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "NCNpp", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "Iaz", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "NNNNNNNp", false)
}

func TestSample5(t *testing.T) {
	runSample(t, "IIqrIIpqIpr", true)
}

func TestSample6(t *testing.T) {
	runSample(t, "Ipp", true)
}

func TestSample7(t *testing.T) {
	runSample(t, "Ezz", true)
}
