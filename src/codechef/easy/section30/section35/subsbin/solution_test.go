package main

import "testing"

func runSample(t *testing.T, s string, expect_length int, expect_cnt int) {
	res, ops := solve(s)

	if res != expect_length || len(ops) != expect_cnt {
		t.Fatalf("Sample expect length %d %d, but got %d %d", expect_length, expect_cnt, res, len(ops))
	}

	buf := []byte(s)
	n := len(s)
	for _, op := range ops {
		l, r, x := op[0], op[1], op[2]
		l--
		r--
		var cnt int
		for i := l; i <= r; i++ {
			if buf[i] == '1' {
				cnt++
			} else {
				cnt--
			}
		}
		if cnt != 0 {
			t.Fatalf("sub string %s, not having equal 0/1", string(buf[l:r+1]))
		}
		if x == 0 {
			buf[l] = '0'
		} else {
			buf[l] = '1'
		}
		copy(buf[l+1:], buf[r+1:])
		n -= (r - l + 1 - 1)
		buf = buf[:n]
	}

	if len(buf) != expect_length {
		t.Fatalf("final string %s, not correct", string(buf))
	}
}

func TestSample1(t *testing.T) {
	s := "11000"
	expect_length := 1
	expect_cnt := 2
	runSample(t, s, expect_length, expect_cnt)
}

func TestSample2(t *testing.T) {
	s := "1100"
	expect_length := 1
	expect_cnt := 1
	runSample(t, s, expect_length, expect_cnt)
}



func TestSample3(t *testing.T) {
	s := "1"
	expect_length := 1
	expect_cnt := 0
	runSample(t, s, expect_length, expect_cnt)
}

func TestSample4(t *testing.T) {
	s := "0000011000"
	expect_length := 1
	expect_cnt := 7
	runSample(t, s, expect_length, expect_cnt)
}
