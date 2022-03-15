package d

import (
	"testing"
)

func runSample(t *testing.T, skills [][]int, expect int) {
	res := coopDevelop(skills)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	skills := [][]int{{1, 2, 3}, {3}, {2, 4}}
	expect := 2
	runSample(t, skills, expect)
}

func TestSample2(t *testing.T) {
	skills := [][]int{{3}, {6}}
	expect := 1
	runSample(t, skills, expect)
}
