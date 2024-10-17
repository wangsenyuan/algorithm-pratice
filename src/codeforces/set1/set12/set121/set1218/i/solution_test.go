package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, src []string, tgt []string, bar string, expect bool) {
	ok, res := solve(src, tgt, bar)

	if expect != ok {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}
	n := len(src)
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = []byte(src[i])
	}

	processRow := func(r int) {
		for i := 0; i < n; i++ {
			if bar[i] == '0' {
				continue
			}
			if buf[r][i] == '1' {
				buf[r][i] = '0'
			} else {
				buf[i][i] = '1'
			}
		}
	}

	processCol := func(c int) {
		for i := 0; i < n; i++ {
			if bar[i] == '0' {
				continue
			}
			if buf[i][c] == '1' {
				buf[i][c] = '0'
			} else {
				buf[i][c] = '1'
			}
		}
	}

	for _, cur := range res {
		var id int
		readInt([]byte(cur), 4, &id)

		if strings.HasPrefix(cur, "row") {
			processRow(id)
		} else {
			processCol(id)
		}
	}

	for i := 0; i < n; i++ {
		if string(buf[i]) != tgt[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	src := []string{
		"11",
		"11",
	}
	tgt := []string{
		"00",
		"01",
	}
	bar := "11"
	expect := false
	runSample(t, src, tgt, bar, expect)
}

func TestSample2(t *testing.T) {
	src := []string{
		"10",
		"00",
	}
	tgt := []string{
		"00",
		"00",
	}
	bar := "10"
	expect := true
	runSample(t, src, tgt, bar, expect)
}

func TestSample3(t *testing.T) {
	src := []string{
		"110",
		"011",
		"100",
	}
	tgt := []string{
		"100",
		"011",
		"100",
	}
	bar := "100"
	expect := true
	runSample(t, src, tgt, bar, expect)
}
