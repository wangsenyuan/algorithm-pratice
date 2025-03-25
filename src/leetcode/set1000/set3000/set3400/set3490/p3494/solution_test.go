package p3494

import (
	"testing"
)

func runSample(t *testing.T, skill []int, mana []int, expect int64) {
	res := minTime(skill, mana)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	skill := []int{1, 5, 2, 4}
	mana := []int{5, 1, 4, 2}
	expect := int64(110)
	runSample(t, skill, mana, expect)
}

func TestSample2(t *testing.T) {
	skill := []int{1, 1, 1}
	mana := []int{1, 1, 1}
	expect := int64(5)
	runSample(t, skill, mana, expect)
}
