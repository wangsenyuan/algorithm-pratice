package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	s := "Abo"
	expect := []int{3, 1}
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := "OEIS"
	expect := []int{3, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "auBAAbeelii"
	expect := []int{9, 3}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "AaaBRAaaCAaaDAaaBRAaa"
	expect := []int{18, 4}
	runSample(t, s, expect)
}