package p3316

import "testing"

func runSample(t *testing.T, source string, pattern string, target []int, expect int) {
	res := maxRemovals(source, pattern, target)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	source := "abbaa"
	pattern := "aba"
	targetIndices := []int{0, 1, 2}
	expect := 1
	runSample(t, source, pattern, targetIndices, expect)
}

func TestSample2(t *testing.T) {
	source := "bcda"
	pattern := "d"
	targetIndices := []int{0, 3}
	expect := 2
	runSample(t, source, pattern, targetIndices, expect)
}

func TestSample3(t *testing.T) {
	source := "dda"
	pattern := "dda"
	targetIndices := []int{0, 1, 2}
	expect := 0
	runSample(t, source, pattern, targetIndices, expect)
}

func TestSample4(t *testing.T) {
	source := "yeyeykyded"
	pattern := "yeyyd"
	targetIndices := []int{0, 2, 3, 4}
	expect := 2
	runSample(t, source, pattern, targetIndices, expect)
}
