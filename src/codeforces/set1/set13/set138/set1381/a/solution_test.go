package main

import "testing"

func runSample(t *testing.T, a, b string) {
	res := solve(a, b)

	n := len(a)
	if len(res) > 2*n {
		t.Fatalf("Sample result %v, took too much steps", res)
	}

	buf := make([]int, n)
	for i := 0; i < n; i++ {
		buf[i] = int(a[i] - '0')
	}

	for _, i := range res {
		i--
		flip(buf[:i+1])
		reverse(buf[:i+1])
	}

	for i := 0; i < n; i++ {
		if int(b[i]-'0') != buf[i] {
			t.Fatalf("Sample result %v got %v, not expected %s", res, buf, b)
		}
	}
}

func flip(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] ^= 1
	}
}


func TestSample1(t *testing.T) {
	a := "01"
	b := "10"
	runSample(t, a, b)
}

func TestSample2(t *testing.T) {
	a := "01011"
	b := "11100"
	runSample(t, a, b)
}

func TestSample3(t *testing.T) {
	a := "0110011011"
	b := "1000110100"
	runSample(t, a, b)
}

func TestSample4(t *testing.T) {
	a := "0101010101"
	b := "1010101010"
	runSample(t, a, b)
}

func TestSample5(t *testing.T) {
	a := "0"
	b := "1"
	runSample(t, a, b)
}
