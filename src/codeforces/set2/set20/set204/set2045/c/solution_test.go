package main

import "testing"

func runSample(t *testing.T, s string, x string, expect string) {
	res := solve(s, x)

	if res == expect {
		return
	}
	if expect == "-1" || res == "-1" || len(expect) != len(res) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	n := len(x)
	m := len(res)
	var cnt int
	for i := 1; i < m; i++ {
		if s[:i] == res[:i] && res[i:] == x[n-(m-i):] {
			cnt++
		}
	}
	if cnt <= 1 {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "sarana"
	x := "olahraga"
	expect := "saga"
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "berhiber"
	x := "wortelhijau"
	expect := "berhijau"
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "icpc"
	x := "icpc"
	expect := "icpc"
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "icpc"
	x := "jakarta"
	expect := "-1"
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "psvpwowlqqjlhxbzhjmuvnmokyvbicq"
	x := "nphnnrnpxoojemjrzdrauykk"
	expect := "psvpwojemjrzdrauykk"
	runSample(t, s, x, expect)
}
