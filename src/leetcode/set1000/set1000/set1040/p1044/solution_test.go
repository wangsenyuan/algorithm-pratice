package p1044

import "testing"

func runSample(t *testing.T, str string, expect string) {
	res := longestDupSubstring(str)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", str, expect, res)
	}
}


func TestSample1(t *testing.T) {
	runSample(t, "banana", "ana")
}


func TestSample2(t *testing.T) {
	runSample(t, "abcd", "")
}