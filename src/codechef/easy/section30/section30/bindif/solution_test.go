package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, cnt int) {
	a, b := solve(len(s), s)
	var cnt2 int
	for i := 0; i < len(a); i++ {
		cnt2 += int(a[i] - '0')
	}

	for i := 0; i < len(b); i++ {
		cnt2 += int(b[i] - '0')
	}
	if cnt2 != cnt {
		t.Errorf("Sample expect %d, but got %d", cnt, cnt2)
	} else {
		var carray int
		m := len(s)
		n := len(b)
		var tmp bytes.Buffer
		for i, j := m-1, n-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
			var x, y int
			if i >= 0 {
				x = int(s[i] - '0')
			}
			if j >= 0 {
				y = int(b[j] - '0')
			}
			z := x + y + carray
			tmp.WriteByte(byte('0' + (z % 2)))
			carray = z / 2
		}
		if carray > 0 {
			tmp.WriteByte('1')
		}
		buf := tmp.Bytes()
		reverse(buf)
		if string(buf) != a {
			t.Errorf("Sample result %s %s, not correct", a, b)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "1110000"
	runSample(t, s, 2)
}


func TestSample2(t *testing.T) {
	s := "10000"
	runSample(t, s, 1)
}
