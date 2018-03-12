package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve([]byte(s))

	if res != expect {
		t.Errorf("sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "808", "818")
}

func TestSample2(t *testing.T) {
	runSample(t, "2133", "2222")
}

func TestSample3(t *testing.T) {
	runSample(t, "99", "101")
}

func TestSample4(t *testing.T) {
	runSample(t, "1999", "2002")
}
