package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `MIKE:MAX.,ARTEM:MIKE..,DMITRY:DMITRY.,DMITRY...`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `A:A..`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `A:C:C:C:C.....`
	expect := 6
	runSample(t, s, expect)
}
