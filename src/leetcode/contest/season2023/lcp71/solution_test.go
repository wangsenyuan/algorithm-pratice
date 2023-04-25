package lcp71

import "testing"

func runSample(t *testing.T, shape []string, expect int) {
	res := reservoir(shape)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	shape := []string{
		"....rl", "l.lr.r", ".l..r.", "..lr..",
	}
	expect := 18
	runSample(t, shape, expect)
}

func TestSample2(t *testing.T) {
	shape := []string{
		".rlrlrlrl", "ll..rl..r", ".llrrllrr", "..lr..lr.",
	}
	expect := 18
	runSample(t, shape, expect)
}

func TestSample3(t *testing.T) {
	shape := []string{
		"l..r", ".lr.",
	}
	expect := 8
	runSample(t, shape, expect)
}

func TestSample4(t *testing.T) {
	shape := []string{
		"lr",
	}
	expect := 2
	runSample(t, shape, expect)
}

func TestSample5(t *testing.T) {
	shape := []string{
		"l....r", ".l..r.", "..lr..",
	}
	expect := 18
	runSample(t, shape, expect)
}
