package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var buf bytes.Buffer
	for _, x := range res {
		if !valid(x) {
			t.Fatalf("Sample result %v, got invalid split %s", res, x)
		}
		buf.WriteString(x)
	}

	ss := buf.String()

	if ss != s {
		t.Fatalf("Sample result %v, got different concatation %s", res, ss)
	}
}

func valid(s string) bool {
	n := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return i > 0 && i <= 8 && (n-(i+1)) > 0 && (n-(i+1)) <= 3
		}
	}
	return false
}

func TestSample1(t *testing.T) {
	s := "read.meexample.txtb.cpp"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "read..txt"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "one.one.one.one.one.one.one.one.one.one"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "z"
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "z.txt"
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "z.txtabcdefgh.a"
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "a.b.c"
	expect := false
	runSample(t, s, expect)
}
