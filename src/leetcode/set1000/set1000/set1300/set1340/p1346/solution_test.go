package p1346

import "testing"

func TestCheckIfExists(t *testing.T) {
	if !checkIfExist([]int{10, 2, 5, 3}) {
		t.Errorf("Fail")
	}

	if !checkIfExist([]int{7, 1, 14, 11}) {
		t.Errorf("Fail")
	}

	if checkIfExist([]int{3, 1, 7, 11}) {
		t.Errorf("Fail")
	}
}

func TestMinSteps(t *testing.T) {
	if 1 != minSteps("bab", "aba") {
		t.Errorf("Fail")
	}
	if 5 != minSteps("leetcode", "practice") {
		t.Errorf("Fail")
	}
	if 0 != minSteps("anagram", "mangaar") {
		t.Errorf("Fail")
	}

	if 4 != minSteps("friend", "family") {
		t.Errorf("Fail")
	}
}
