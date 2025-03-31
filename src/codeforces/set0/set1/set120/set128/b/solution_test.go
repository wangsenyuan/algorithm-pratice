package main

import (
	"math/rand"
	"testing"
	"time"
)

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample %s %d expect %s, but got %s", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aa"
	k := 2
	expect := "a"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	k := 5
	// a, ab, abc, b, bc, c
	expect := "bc"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "abab"
	k := 7
	expect := "b"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "abcd"
	k := 4*5/2 - 1
	expect := "cd"
	runSample(t, s, k, expect)
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

func TestSample5(t *testing.T) {
	for range 3 {
		n := 6
		s := randomString(6)
		t.Logf("will test %s", s)
		for k := 1; k <= n*(n+1)/2; k++ {
			runSample(t, s, k, bruteForce(s, k))
		}
	}
}

func TestSample6(t *testing.T) {
	s := "bcbjdx"
	k := 3
	expect := bruteForce(s, k)
	runSample(t, s, k, expect)
}

func TestSample7(t *testing.T) {
	s := "lfhhoh"
	k := 2
	expect := bruteForce(s, k)
	runSample(t, s, k, expect)
}

func TestSample8(t *testing.T) {
	s := "oiadec"
	k := 4
	// a ad ade adec
	expect := bruteForce(s, k)
	runSample(t, s, k, expect)
}
