package main

import (
	"math/rand"
	"testing"
	"time"
)

func runSample(t *testing.T, s string, expect [][]int) {

	getLength := func(arr [][]int) int {
		var sum int
		for _, cur := range arr {
			sum += cur[1]
		}
		return sum
	}

	res := solve(s)
	if getLength(res) != getLength(expect) {
		t.Fatalf("sample %s, expect %v, but got %v", s, expect, res)
	}

	var x string
	for _, cur := range res {
		x += s[cur[0]-1 : cur[0]+cur[1]-1]
	}

	if reverse(x) != x {
		t.Fatalf("Sample result %v, getting a non-palindrome %s", res, x)
	}
}

func TestSample1(t *testing.T) {
	s := "abacaba"
	expect := [][]int{{1, 7}}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "axbya"
	expect := [][]int{{1, 1}, {2, 1}, {5, 1}}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "xabyczba"
	expect := [][]int{{2, 2}, {4, 1}, {7, 2}}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "aaabaaaaaa"
	expect := [][]int{{1, 3}, {5, 3}, {8, 3}}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "aaaaaabaaa"
	expect := [][]int{{1, 3}, {4, 3}, {8, 3}}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	x := randomString(1000)
	y := reverse(x[:len(x)-1])
	x += y
	a := x[:100]
	b := x[100:101]
	c := x[101:]
	s := randomString(10) + a + randomString(10) + b + randomString(11) + c
	expect := bruteForce(s)
	runSample(t, s, expect)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
