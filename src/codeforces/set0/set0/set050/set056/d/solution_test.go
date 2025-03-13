package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	s, x, ops, pos, chs := process(bufio.NewReader(strings.NewReader(s)))

	if len(ops) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, ops)
	}
	buf := make([]byte, len(s)+len(x))
	copy(buf, s)
	n := len(s)
	for i := range expect {
		op := ops[i]
		if op == "INSERT" {
			j := pos[i] - 1
			if j < 0 || j > n {
				t.Fatalf("Sample result not correct, the pos %d is out of bounds", j)
			}
			copy(buf[j+1:], buf[j:])
			buf[j] = chs[i]
			n++
		} else if op == "REPLACE" {
			j := pos[i] - 1
			if j < 0 || j >= n {
				t.Fatalf("Sample result not correct, the pos %d is out of bounds", j)
			}
			buf[j] = chs[i]
		} else {
			// DELETE
			j := pos[i] - 1
			if j < 0 || j >= n {
				t.Fatalf("Sample result not correct, the pos %d is out of bounds", j)
			}
			copy(buf[j:], buf[j+1:])
			n--
		}
	}
	if n != len(x) || string(buf[:n]) != x {
		t.Fatalf("Sample result %s, not expected %s", string(buf[:n]), x)
	}
}

func TestSample1(t *testing.T) {
	s := `ABA
ABBBA
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `ACCEPTED
WRONGANSWER
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `TJHGFKKCDOHRNAXZROCWIYFF
ZBWUEHEVEOUATECAGLZIQMUDXEMHRSOZMAUJRWLQMPPZOUMXHAMWUGEDIKVKBLVMXWUOFMPAFDPRBCFTEWOULCZWRQHCTBTBXRHHODWBCXWIMNCNEXOSKSUJLISGCLLLXOKRSBNOZTHAJNNLILYFFMSYKOFPTXRNEFBSOUHFOLTIQAINRPXWRQ
`
	expect := 164
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `ABA
B
`
	expect := 2
	runSample(t, s, expect)
}
