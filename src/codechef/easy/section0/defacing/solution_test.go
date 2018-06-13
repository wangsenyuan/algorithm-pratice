package main

import "testing"

func runSample(t *testing.T, s []byte, m []byte, expect int) {
	res := solve(s, m)
	if res != expect {
		t.Errorf("solve %s, %s should give %d, but got %d", s, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []byte("25")
	m := []byte("100")
	runSample(t, s, m, 89)
}

func TestSample2(t *testing.T) {
	s := []byte("77")
	m := []byte("100")
	runSample(t, s, m, 100)
}

func TestSample3(t *testing.T) {
	s := []byte("0")
	m := []byte("100")
	runSample(t, s, m, 100)
}

func TestSample4(t *testing.T) {
	s := []byte("8")
	m := []byte("13")
	runSample(t, s, m, 8)
}

func TestSample5(t *testing.T) {
	s := []byte("8")
	m := []byte("24")
	runSample(t, s, m, 18)
}

func TestSample6(t *testing.T) {
	s := []byte("0")
	m := []byte("0")
	runSample(t, s, m, 0)
}

func TestSample7(t *testing.T) {
	s := []byte("19")
	m := []byte("38")
	runSample(t, s, m, 38)
}

func TestSample8(t *testing.T) {
	s := []byte("89375")
	m := []byte("9247529")
	runSample(t, s, m, 9189999)
}

func TestSample9(t *testing.T) {
	s := []byte("804276")
	m := []byte("2857282")
	runSample(t, s, m, 2809898)
}

func TestSample10(t *testing.T) {
	s := []byte("0")
	m := []byte("20130120")
	runSample(t, s, m, 20130120)
}

func TestSample11(t *testing.T) {
	s := []byte("3284709")
	m := []byte("20130120")
	runSample(t, s, m, 19889989)
}

func TestSample12(t *testing.T) {
	s := []byte("1")
	m := []byte("6")
	runSample(t, s, m, 4)
}
