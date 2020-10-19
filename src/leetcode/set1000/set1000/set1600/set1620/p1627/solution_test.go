package p1627

import "testing"

func runSample(t *testing.T, scores []int, ages []int, expect int) {
	res := bestTeamScore(scores, ages)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", scores, ages, expect, res)
	}
}

func TestSample1(t *testing.T) {
	scores := []int{1, 3, 5, 10, 15}
	ages := []int{1, 2, 3, 4, 5}
	expect := 34
	runSample(t, scores, ages, expect)
}

func TestSample2(t *testing.T) {
	scores := []int{4, 5, 6, 5}
	ages := []int{2, 1, 2, 1}
	expect := 16
	runSample(t, scores, ages, expect)
}

func TestSample3(t *testing.T) {
	scores := []int{1, 2, 3, 5}
	ages := []int{8, 9, 10, 1}
	expect := 6
	runSample(t, scores, ages, expect)
}

func TestSample4(t *testing.T) {
	scores := []int{9, 2, 8, 8, 2}
	ages := []int{4, 1, 3, 3, 5}
	expect := 27
	runSample(t, scores, ages, expect)
}
