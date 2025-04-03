package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res[0] != expect[0] || res[1] != expect[1] {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 70 100 100 25`
	expect := []int{99, 33}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `300 500 1000 1000 300`
	expect := []int{1000, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `143 456 110 117 273`
	expect := []int{76, 54}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `100 110 2 2 109`
	expect := []int{0, 2}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 1 100 100 1`
	expect := []int{100, 100}
	runSample(t, s, expect)
}
