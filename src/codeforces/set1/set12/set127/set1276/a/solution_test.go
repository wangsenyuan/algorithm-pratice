package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	var buf bytes.Buffer

	for i, j := 1, 0; i <= len(s); i++ {
		if j < len(res) && res[j] == i {
			j++
			continue
		}
		buf.WriteByte(s[i-1])
	}

	x := buf.String()

	for i := 0; i+3 <= len(x); i++ {
		if x[i:i+3] == "one" || x[i:i+3] == "two" {
			t.Fatalf("Sample result %v, got %s, has %s at pos %d", res, x, x[i:i+3], i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "onetwone"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "testme"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "oneoneone"
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "onetwonetwooneooonetwooo"
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "onnne"
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "oneeeee"
	expect := 1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "oneeeeeeetwooooo"
	expect := 2
	runSample(t, s, expect)
}
