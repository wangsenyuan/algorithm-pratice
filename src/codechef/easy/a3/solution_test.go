package main

import "testing"

func TestSample1(t *testing.T) {
	hints := [][]byte{
		[]byte("< 100 No"),
		[]byte("> 100 No"),
	}
	n := 2
	res := solve(n, hints)
	if res != 0 {
		t.Errorf("this sample should give answer 0, but get %d", res)
	}
}

func TestSample2(t *testing.T) {
	hints := [][]byte{
		[]byte("< 2 Yes"),
		[]byte("> 4 Yes"),
		[]byte("= 3 No"),
	}
	n := 3
	res := solve(n, hints)
	if res != 1 {
		t.Errorf("this sample should give answer 1, but get %d", res)
	}
}

func TestSample3(t *testing.T) {
	hints := [][]byte{
		[]byte("< 2 Yes"),
		[]byte("> 1 Yes"),
		[]byte("= 1 Yes"),
		[]byte("= 1 Yes"),
		[]byte("> 1 Yes"),
		[]byte("= 1 Yes"),
	}
	n := 6
	res := solve(n, hints)
	if res != 2 {
		t.Errorf("this sample should give answer 2, but get %d", res)
	}
}
