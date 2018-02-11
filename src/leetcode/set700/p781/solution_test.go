package p781

import "testing"

func runSample(t *testing.T, answers []int, expect int) {
	res := numRabbits(answers)

	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", answers, expect, res)
	}
}

func TestSample1(t *testing.T) {
	answers := []int{1, 1, 2}
	runSample(t, answers, 5)
}

func TestSample2(t *testing.T) {
	answers := []int{1, 0, 1, 0, 0}
	runSample(t, answers, 5)
}
