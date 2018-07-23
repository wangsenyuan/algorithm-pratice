package p806

import "testing"

func runSample(t *testing.T, s string, widths []int, expect []int) {
	res := numberOfLines(widths, s)
	if res[0] != expect[0] || res[1] != expect[1] {
		t.Errorf("sample %s %v, expect %v, but got %v", s, widths, expect, res)
	}
}

func TestSample1(t *testing.T) {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "abcdefghijklmnopqrstuvwxyz"
	expect := []int{3, 60}
	runSample(t, S, widths, expect)
}

func TestSample2(t *testing.T) {
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "bbbcccdddaaa"
	expect := []int{2, 4}
	runSample(t, S, widths, expect)
}
