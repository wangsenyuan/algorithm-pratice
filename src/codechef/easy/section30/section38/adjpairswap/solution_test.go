package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	ok, swaps := solve(s)

	if ok != expect {
		t.Errorf("Sample expect %t, but got %t", expect, ok)
		return
	}
	if ok {
		buf := []byte(s)
		for _, swap := range swaps {
			i, j := swap[0], swap[1]
			i--
			j--
			buf[i], buf[j] = buf[j], buf[i]
			buf[i+1], buf[j+1] = buf[j+1], buf[i+1]
		}

		var x int
		for x < len(buf) && buf[x] == '0' {
			x++
		}
		if x < len(buf) {
			for x < len(buf) && buf[x] == '1' {
				x++
			}
			if x < len(buf) {
				t.Errorf("Sample result %s after swaps %v, not correct", string(buf), swaps)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := "0110111"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0 1 0 1 0 0 1 1"
	s = strings.ReplaceAll(s, " ", "")
	expect := true
	runSample(t, s, expect)
}
