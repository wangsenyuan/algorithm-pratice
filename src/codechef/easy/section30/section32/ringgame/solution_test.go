package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, A string, expect string) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "101"
	expect := BOB
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := "1 1 0 0"
	A = strings.ReplaceAll(A, " ", "")
	expect := ALICE
	runSample(t, A, expect)
}
