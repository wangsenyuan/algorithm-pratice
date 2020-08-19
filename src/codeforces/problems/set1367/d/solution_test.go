package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, m int, B []int) {
	C := make([]int, m)
	copy(C, B)
	res := solve(s, m, C)

	b := check(res)

	if !reflect.DeepEqual(b, B) {
		t.Errorf("Sample %s %d %v, result %s not correct", s, m, B, res)
	}
}

func check(s string) []int {

	m := len(s)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		if s[i] < 'a' || s[i] > 'z' {
			res[i] = -1
			return res
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if s[j] > s[i] {
				res[i] += abs(i - j)
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	s := "abac"
	m := 3
	B := []int{2, 1, 0}
	runSample(t, s, m, B)
}

func TestSample2(t *testing.T) {
	s := "abc"
	m := 1
	B := []int{0}
	runSample(t, s, m, B)
}

func TestSample3(t *testing.T) {
	s := "abba"
	m := 3
	B := []int{1, 0, 1}
	runSample(t, s, m, B)
}

func TestSample4(t *testing.T) {
	s := "ecoosdcefr"
	m := 10
	B := []int{38, 13, 24, 14, 11, 5, 3, 24, 17, 0}
	runSample(t, s, m, B)
}
