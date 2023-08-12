package main

import "testing"

func runSample(t *testing.T, h int, m int, s string, expect string) {
	res := solve(h, m, s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h, m := 24, 60
	s := "12:21"
	expect := "12:21"
	runSample(t, h, m, s, expect)
}

func TestSample2(t *testing.T) {
	h, m := 24, 60
	s := "23:59"
	expect := "00:00"
	runSample(t, h, m, s, expect)
}

func TestSample3(t *testing.T) {
	h, m := 90, 80
	s := "52:26"
	expect := "52:28"
	runSample(t, h, m, s, expect)
}

func TestSample4(t *testing.T) {
	h, m := 1, 100
	s := "00:01"
	expect := "00:00"
	runSample(t, h, m, s, expect)
}

func TestSample5(t *testing.T) {
	h, m := 10, 10
	s := "04:04"
	expect := "00:00"
	runSample(t, h, m, s, expect)
}

func TestSample6(t *testing.T) {
	h, m := 100, 100
	s := "02:82"
	expect := "02:82"
	runSample(t, h, m, s, expect)
}

func TestSample7(t *testing.T) {
	h, m := 51, 3
	s := "30:01"
	expect := "50:00"
	runSample(t, h, m, s, expect)
}

func TestSample8(t *testing.T) {
	h, m := 35, 7
	s := "10:02"
	expect := "10:05"
	runSample(t, h, m, s, expect)
}
