package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, expect []string, ok bool) {
	flag, ops := solve(expect)

	if flag != ok {
		t.Fatalf("Sample expect %t, but got %t", ok, flag)
	}
	if !ok {
		return
	}
	res := paint(ops, len(expect), len(expect[0]))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample after %v, got %v, not expected %v", ops, res, expect)
	}
}

func paint(ops [][]int, n int, m int) []string {
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			buf[i][j] = '0'
		}
	}

	pt := func(a, b, c, d int) {
		for i := a; i <= c; i++ {
			for j := b; j <= d; j++ {
				dx := i - a
				dy := j - b
				if (dx+dy)%2 == 0 {
					buf[i][j] = '0'
				} else {
					buf[i][j] = '1'
				}
			}
		}
	}

	for _, op := range ops {
		a, b, c, d := op[0], op[1], op[2], op[3]
		pt(a-1, b-1, c-1, d-1)
	}

	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = string(buf[i])
	}
	return res
}

func TestSample1(t *testing.T) {
	expect := []string{
		"01000",
		"10100",
		"01010",
		"00110",
	}
	ok := true
	runSample(t, expect, ok)
}

func TestSample2(t *testing.T) {
	expect := []string{
		"001",
		"010",
	}
	ok := true
	runSample(t, expect, ok)
}

func TestSample3(t *testing.T) {
	expect := []string{
		"110",
		"101",
		"000",
	}
	ok := false
	runSample(t, expect, ok)
}
