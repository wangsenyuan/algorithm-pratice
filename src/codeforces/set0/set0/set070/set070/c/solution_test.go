package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if reflect.DeepEqual(expect, res) {
		return
	}
	if len(expect) == 0 || len(res) == 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if expect[0]*expect[1] != res[0]*res[1] {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	s := "2 2 1"
	expect := []int{1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "132 10 35"
	expect := []int{7, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "5 18 1000"
	runSample(t, s, nil)
}

func TestSample4(t *testing.T) {
	s := "48 132 235"
	expect := []int{22, 111}
	runSample(t, s, expect)
}
