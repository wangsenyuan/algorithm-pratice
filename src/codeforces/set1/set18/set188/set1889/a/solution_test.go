package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ok, res := solve(s)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}
	if !expect {
		return
	}
	buf := []byte(s)

	for _, i := range res {
		tmp := make([]byte, len(buf)+2)
		if i == 0 {
			tmp[0] = '0'
			tmp[1] = '1'
			copy(tmp[2:], buf)
		} else {
			copy(tmp[:i], buf[:i])
			tmp[i] = '0'
			tmp[i+1] = '1'
			copy(tmp[i+2:], buf[i:])
		}

		buf = tmp
	}

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		if buf[i] == buf[j] {
			t.Fatalf("Sample result %v, getting wrong result %s", res, string(buf))
		}
	}
}

func TestSample1(t *testing.T) {
	s := "1111"
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "001110"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0111001100"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0010"
	expect := false
	runSample(t, s, expect)
}
