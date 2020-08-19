package main

import "testing"

func runSample(t *testing.T, S, P string, expect string) {
	res := solve(S, P)

	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", S, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "akramkeeanany"
	P := "aka"
	expect := "aaakaeekmnnry"
	runSample(t, S, P, expect)
}

func TestSample2(t *testing.T) {
	S := "supahotboy"
	P := "bohoty"
	expect := "abohotypsu"
	runSample(t, S, P, expect)
}

func TestSample3(t *testing.T) {
	S := "daehabshatorawy"
	P := "badawy"
	expect := "aabadawyehhorst"
	runSample(t, S, P, expect)
}
