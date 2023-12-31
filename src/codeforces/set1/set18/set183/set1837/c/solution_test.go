package main

import "testing"

func runSample(t *testing.T, p string, expect string) {
	res := solve(p)

	if cost(res) != cost(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func cost(s string) int {
	var res int

	for i := 0; i+1 < len(s); i++ {
		if s[i] == '1' && s[i+1] == '0' {
			res++
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	p := "??01?"
	expect := "00011"
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := "10100"
	expect := "10100"
	runSample(t, p, expect)
}
func TestSample3(t *testing.T) {
	p := "1??10?"
	expect := "111101"
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := "0?1?10?10"
	expect := "011110010"
	runSample(t, p, expect)
}
