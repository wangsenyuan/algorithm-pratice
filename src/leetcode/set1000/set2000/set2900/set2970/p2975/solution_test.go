package p2975

import "testing"

func runSample(t *testing.T, source string, target string, original []string, changed []string, cost []int, expect int) {
	res := minimumCost(source, target, original, changed, cost)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	source := "abcdefgh"
	target := "acdeeghh"
	original := []string{"bcd", "fgh", "thh"}
	changed := []string{"cde", "thh", "ghh"}
	cost := []int{1, 3, 5}
	expect := 9
	runSample(t, source, target, original, changed, cost, expect)
}

func TestSample2(t *testing.T) {
	source := "abcd"
	target := "acbe"
	original := []string{"a", "b", "c", "c", "e", "d"}
	changed := []string{"b", "c", "b", "e", "b", "e"}
	cost := []int{2, 5, 5, 1, 2, 20}
	expect := 28
	runSample(t, source, target, original, changed, cost, expect)
}

func TestSample3(t *testing.T) {
	source := "abcdefgh"
	target := "addddddd"
	original := []string{"bcd", "defgh"}
	changed := []string{"ddd", "ddddd"}
	cost := []int{100, 1578}
	expect := -1
	runSample(t, source, target, original, changed, cost, expect)
}
