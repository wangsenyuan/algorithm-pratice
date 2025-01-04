package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res, pairs := process(bufio.NewReader(strings.NewReader(s)))

	if len(expect) != len(res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	var arr []int
	for _, i := range res {
		p := pairs[i-1]
		arr = append(arr, p[0], p[1])
	}

	if arr[0] == arr[1] {
		t.Fatalf("Sample result %v, got wrong sequence %v", res, arr)
	}
	for i := 2; i < len(res); i += 2 {
		if arr[i] == arr[i+1] || (arr[i] > arr[i+1]) != (arr[0] > arr[1]) {
			t.Fatalf("Sample result %v, got wrong sequence %v", res, arr)
		}
		if arr[i-1] == arr[i] || arr[i-1] < arr[i] != (arr[0] > arr[1]) {
			t.Fatalf("Sample result %v, got wrong sequence %v", res, arr)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 7
6 4
2 10
9 8
3 5
`
	expect := []int{1, 5, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
5 4
3 2
6 1
`
	expect := []int{3, 2, 1}
	runSample(t, s, expect)
}