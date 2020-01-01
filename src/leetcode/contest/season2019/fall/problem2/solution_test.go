package problem2

import "testing"

func runSample(t *testing.T, cont []int, expect []int) {
	res := fraction(cont)

	if res[0] != expect[0] && res[1] != expect[1] {
		t.Errorf("Sample %v, expect %v, but got %v", cont, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cont := []int{3, 2, 0, 2}
	expect := []int{13, 4}
	runSample(t, cont, expect)
}
func TestSample2(t *testing.T) {
	cont := []int{0, 0, 3}
	expect := []int{3, 1}
	runSample(t, cont, expect)
}

func TestSample3(t *testing.T) {
	cont := []int{0, 2}
	expect := []int{1, 2}
	runSample(t, cont, expect)
}

func TestSample4(t *testing.T) {
	cont := []int{2}
	expect := []int{2, 1}
	runSample(t, cont, expect)
}
