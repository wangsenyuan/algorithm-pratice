package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ok, b, p := solve(s)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}
	buf := []byte(s)
	if len(b) > 0 {
		x := buf[b[len(b)-1]-1]
		for j := len(b) - 1; j > 0; j-- {
			buf[b[j]-1] = buf[b[j-1]-1]
		}
		buf[b[0]-1] = x
	}
	var u, v []byte
	for i, j := 0, 0; i < len(s); i++ {
		if j < len(p) && p[j]-1 == i {
			u = append(u, buf[i])
			j++
		} else {
			v = append(v, buf[i])
		}
	}

	if string(u) != string(v) {
		t.Fatalf("Sample result %v, %v not correct", b, p)
	}
}

func TestSample1(t *testing.T) {
	s := "100010"
	expect := true
	runSample(t, s, expect)
}
