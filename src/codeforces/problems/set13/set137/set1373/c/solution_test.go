package main

import "testing"

func runSample(t *testing.T, s string) {
	res := solve(len(s), []byte(s))
	expect := simulate(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func simulate(s string) int64 {
	var init int
	var res int64
	for {
		cur := init
		ok := true
		for i := 0; i < len(s); i++ {
			res++
			if s[i] == '+' {
				cur++
			} else {
				cur--
			}
			if cur < 0 {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		init++
	}
	return res
}

func TestSample1(t *testing.T) {
	runSample(t, "--+-")
}

func TestSample2(t *testing.T) {
	runSample(t, "---")
}

func TestSample3(t *testing.T) {
	runSample(t, "++--+-")
}
