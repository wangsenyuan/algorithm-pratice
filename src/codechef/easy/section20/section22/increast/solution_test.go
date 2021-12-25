package main

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := solve(len(S), S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "aba"
	expect := "aab"
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "abcd"
	expect := "abcd"
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "cbcdbef"
	expect := "bbccdef"
	runSample(t, S, expect)
}


func TestSample4(t *testing.T) {
	S := "fabcdac"
	expect := "aacfbcd"
	runSample(t, S, expect)
}

func TestSample5(t *testing.T) {
	S := "aadcacd"
	expect := "aaacdcd"
	runSample(t, S, expect)
}


func TestSample6(t *testing.T) {
	S := "ababacbcbc"
	expect := "aaabbbbccc"
	runSample(t, S, expect)
}

func TestSample7(t *testing.T) {
	S := "pqpqrqs"
	expect := "ppqqqrs"
	runSample(t, S, expect)
}


func TestSample8(t *testing.T) {
	S := "aadeacd"
	expect := "aaacdde"
	runSample(t, S, expect)
}

func TestSample9(t *testing.T) {
	S := "aadcacd"
	expect := "aaacdcd"
	runSample(t, S, expect)
}
