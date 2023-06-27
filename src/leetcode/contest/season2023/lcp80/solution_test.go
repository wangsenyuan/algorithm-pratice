package lcp80

import "testing"

func runSample(t *testing.T, parents []int, expect string) {
	res := evolutionaryRecord(parents)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	parents := []int{-1, 0, 0, 2}
	expect := "00110"
	runSample(t, parents, expect)
}

func TestSample2(t *testing.T) {
	parents := []int{-1, 0, 0, 1, 2, 2}
	expect := "00101100"
	runSample(t, parents, expect)
}

func TestSample3(t *testing.T) {
	parents := []int{-1, 0, 1, 1, 3, 0, 1, 1, 6, 0, 9, 6, 7, 4, 12, 1, 3, 5}
	expect := "00001101100011100101101011001100"
	runSample(t, parents, expect)
}
